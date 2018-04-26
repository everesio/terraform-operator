package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"text/tabwriter"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

const apiVersion = "v1"

var extraStructBuffer = map[string]string{}

func main() {
	prov := aws.Provider().(*schema.Provider)
	for k, v := range prov.ResourcesMap {
		err := generateTypeFile(titleStr(k), fmt.Sprintf("../pkg/apis/aws/v1/%s.go", k), v.Schema)
		if err != nil {
			panic(err)
		}

	}

	err := generateRegisterFile("../pkg/apis/aws/v1/register.go", prov.ResourcesMap)
	if err != nil {
		panic(err)
	}
}

func generateTypeFile(name, path string, m map[string]*schema.Schema) error {
	b := []byte{}
	bb := bytes.NewBuffer(b)
	w := new(tabwriter.Writer)
	w.Init(bb, 0, 8, 1, '\t', 0)
	fmt.Fprintf(bb, `
package %s

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
%s
%s
%s
`, apiVersion, generateTypeStructString(name), generateSpecStructString(name, m), generateTypeStructStringList(name))

	for _, v := range extraStructBuffer {
		fmt.Fprintf(bb, v)
	}
	extraStructBuffer = map[string]string{}
	w.Flush()
	err := ioutil.WriteFile(path, bb.Bytes(), 0755)
	if err != nil {
		return err
	}

	return nil
}

func generateTypeStructStringList(n string) string {
	tlist := fmt.Sprintf(`
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// %sList is a list of %s resources
type %sList struct {
%s

%s
}`, n, n, n, buildMetaObjects(), buildTypeStructBodyList(n))

	return fmt.Sprintf("%s\n", tlist)
}

func generateTypeStructString(n string) string {
	t := fmt.Sprintf(`
// %s describes a %s resource
type %s struct {
%s

%s
}`, n, n, n, buildMetaObjects(), buildTypeStructBody(n))

	return fmt.Sprintf("%s\n", t)
}

func buildMetaObjects() string {
	return strings.Join([]string{
		"\tmeta_v1.TypeMeta\t`json:\",inline\"`",
		"\tmeta_v1.ObjectMeta\t`json:\"metadata,omitempty\"`",
	}, "\n")
}

func buildTypeStructBody(n string) string {
	return fmt.Sprintf("\tSpec\t%sSpec\t`json:\"spec\"`", n)
}
func buildTypeStructBodyList(n string) string {
	return fmt.Sprintf("\tItems\t[]%s\t`json:\"items\"`", n)
}

func generateSpecStructString(n string, m map[string]*schema.Schema) string {
	spec := fmt.Sprintf(`
// %sSpec is the spec for a %s Resource
type %sSpec struct {
`, n, n, n)
	for k, v := range m {
		l := generateFieldString(k, v)
		if l != "" {
			spec = spec + l + "\n"
		}
	}
	spec = spec + "}"
	return spec
}

func generateFieldString(k string, v *schema.Schema) string {
	name := titleStr(k)
	t := typeSchemaAsString(name, v)
	if v.Computed {
		return ""
	}
	return fmt.Sprintf("\t%s\t%s\t`json:\"%s\"`", name, t, k)
}

func titleStr(s string) string {
	p := strings.Replace(s, "_", " ", -1)
	t := strings.Title(p)
	return strings.Replace(t, " ", "", -1)
}

func typeSchemaAsString(name string, t *schema.Schema) string {
	switch t.Type {
	case schema.TypeString:
		return "string"
	case schema.TypeBool:
		return "bool"
	case schema.TypeFloat:
		return "float64"
	case schema.TypeInt:
		return "int"
	case schema.TypeList:
		return "[]" + typeElemAsString(name, t.Elem)
	case schema.TypeMap:
		return "map[string]" + typeElemAsString(name, t.Elem)
	case schema.TypeSet:
		return typeElemAsString(name, t.Elem)
	}

	return "string"
}

func typeElemAsString(name string, i interface{}) string {
	switch i.(type) {
	case *schema.Schema:
		return typeSchemaAsString(name, i.(*schema.Schema))
	case *schema.Resource:
		// Complex object, shit ourselves.
		generateInlineStruct(name, i.(*schema.Resource))
		return name
	}

	return "string"
}

func generateInlineStruct(name string, r *schema.Resource) {
	spec := fmt.Sprintf(`
// %s is a %s
type %s struct {
`, name, name, name)
	for k, v := range r.Schema {
		l := generateFieldString(k, v)
		if l != "" {
			spec = spec + l + "\n"
		}
	}
	spec = spec + "}\n"
	extraStructBuffer[name] = spec
}

func generateRegisterFile(path string, m map[string]*schema.Resource) error {
	b := []byte{}
	bb := bytes.NewBuffer(b)
	w := new(tabwriter.Writer)
	w.Init(bb, 0, 8, 1, '\t', 0)

	foo := fmt.Sprintf(`
	package %s
	import (
		meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		"k8s.io/apimachinery/pkg/runtime"
		"k8s.io/apimachinery/pkg/runtime/schema"
	
		"github.com/trussle/terraform-operator/pkg/apis/aws"
	)
	
	// GroupVersion is the identifier for the API which includes
	// the name of the group and the version of the API
	var SchemeGroupVersion = schema.GroupVersion{
		Group:   aws.GroupName,
		Version: "%s",
	}
	
	// create a SchemeBuilder which uses functions to add types to
	// the scheme
	var AddToScheme = runtime.NewSchemeBuilder(addKnownTypes).AddToScheme
	
	func Resource(resource string) schema.GroupResource {
		return SchemeGroupVersion.WithResource(resource).GroupResource()
	}
	%s`, apiVersion, apiVersion, generateKnownTypesString(m))
	fmt.Fprintf(bb, foo)
	w.Flush()
	err := ioutil.WriteFile(path, bb.Bytes(), 0755)
	if err != nil {
		return err
	}

	return nil
}

func generateKnownTypesString(m map[string]*schema.Resource) string {
	return fmt.Sprintf(`func addKnownTypes(scheme *runtime.Scheme) error {
		%s
	
		// register the type in the scheme
		meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
		return nil
	}`, generateAddKnownTypesCallString(m))
}

func generateAddKnownTypesCallString(m map[string]*schema.Resource) string {
	o := "schema.AddKnownTypes(\n\t\t\tSchemeGroupVersion,"
	for k := range m {
		o = o + fmt.Sprintf("\n\t\t\t&%s{},", titleStr(k))
		o = o + fmt.Sprintf("\n\t\t\t&%sList{},", titleStr(k))
	}
	o = o + "\n)"
	return o
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SliceUniq(s []string) []string {
	for i := 0; i < len(s); i++ {
		for i2 := i + 1; i2 < len(s); i2++ {
			if s[i] == s[i2] {
				// delete
				s = append(s[:i2], s[i2+1:]...)
				i2--
			}
		}
	}
	return s
}
