package components

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ComponentList_New(t *testing.T) {
	t.Run("From YAML", func(t *testing.T) {
		clList, err := NewComponentList("../test/data/componentlist.yaml")
		require.NoError(t, err)
		verifyComponentList(t, clList)
	})
	t.Run("From JSON", func(t *testing.T) {
		clList, err := NewComponentList("../test/data/componentlist.json")
		require.NoError(t, err)
		verifyComponentList(t, clList)
	})
}

func Test_ComponentList_Remove(t *testing.T) {
	t.Run("Remove Prerequisite", func(t *testing.T) {
		clList, err := NewComponentList("../test/data/componentlist.yaml")
		require.NoError(t, err)
		verifyComponentList(t, clList)
		clList.Remove("prereqcomp1")
		require.Equal(t, 1, len(clList.Prerequisites), "Different amount of prerequisite components")
		require.Equal(t, 3, len(clList.Components), "Different amount of components")
		require.Equal(t, "prereqcomp2", clList.Prerequisites[0].Name)
	})
	t.Run("Remove Component", func(t *testing.T) {
		clList, err := NewComponentList("../test/data/componentlist.yaml")
		require.NoError(t, err)
		verifyComponentList(t, clList)
		clList.Remove("comp2")
		require.Equal(t, 2, len(clList.Prerequisites), "Different amount of prerequisite components")
		require.Equal(t, 2, len(clList.Components), "Different amount of components")
		require.Equal(t, "comp1", clList.Components[0].Name)
		require.Equal(t, "comp3", clList.Components[1].Name)
	})
}

func verifyComponentList(t *testing.T, clList *ComponentList) {
	prereqs := clList.Prerequisites
	comps := clList.Components
	// verify amount of components

	require.Equal(t, 2, len(prereqs), "Different amount of prerequisite components")
	require.Equal(t, 3, len(comps), "Different amount of components")

	// verify names + namespaces of prerequisites
	require.Equal(t, "prereqcomp1", prereqs[0].Name, "Wrong component name")
	require.Equal(t, "prereqns1", prereqs[0].Namespace, "Wrong namespace")
	require.Equal(t, "prereqcomp2", prereqs[1].Name, "Wrong component name")
	require.Equal(t, "testns", prereqs[1].Namespace, "Wrong namespace")

	// verify names + namespaces of components
	require.Equal(t, "comp1", comps[0].Name, "Wrong component name")
	require.Equal(t, "testns", comps[0].Namespace, "Wrong namespace")
	require.Equal(t, "comp2", comps[1].Name, "Wrong component name")
	require.Equal(t, "compns2", comps[1].Namespace, "Wrong namespace")
	require.Equal(t, "comp3", comps[2].Name, "Wrong component name")
	require.Equal(t, "testns", comps[2].Namespace, "Wrong namespace")
}
