package mapping

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/omniscale/imposm3/element"
)

// go test  ./mapping -run TestAdvancedFiltering   -v
func TestAdvancedFiltering(t *testing.T) {

	var configTestMapping *Mapping
	var err error

	configTestMapping, err = NewMapping("./config_test_mapping.yml")
	if err != nil {
		panic(err)
	}

	var wayfilterTests = []struct {
		expected  []int        // expected result for every test table [t0 - tn table]  1=true 0=false  see: config_test_mapping.yml
		inputtags element.Tags // input  osm tags
	}{

		//  expected t0..tn          ||                    inputtags for test
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "0", "boundary": "administrative"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "1", "boundary": "administrative"}},
		{[]int{1, 1, 1, 1, 1, 1, 1}, element.Tags{"admin_level": "2", "boundary": "administrative", "name": "N2"}},

		{[]int{0, 0, 0, 0, 0, 0, 1}, element.Tags{"admin_level": "3", "boundary": "administrative", "name": "N3"}},
		{[]int{1, 1, 1, 1, 1, 1, 1}, element.Tags{"admin_level": "4", "boundary": "administrative", "name": "N4"}},
		{[]int{0, 0, 0, 0, 0, 0, 1}, element.Tags{"admin_level": "5", "boundary": "administrative", "name": "N5"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "6", "boundary": "administrative"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "7", "boundary": "administrative"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "8", "boundary": "administrative"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "X", "boundary": "administrative"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "Y", "boundary": "administrative"}},

		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "0", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "1", "boundary": "maritime"}},
		{[]int{1, 1, 1, 1, 1, 1, 0}, element.Tags{"admin_level": "2", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 1}, element.Tags{"admin_level": "3", "boundary": "maritime", "name": "N3"}},
		{[]int{1, 1, 1, 1, 1, 1, 0}, element.Tags{"admin_level": "4", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "5", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "6", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 1}, element.Tags{"admin_level": "7", "boundary": "maritime", "name": "N7"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "8", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "X", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "Y", "boundary": "maritime"}},

		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "0", "boundary": "political"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "1", "boundary": "political", "name": "N1"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "2", "boundary": "political"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "3", "boundary": "political", "name": "N3"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "4", "boundary": "political"}},

		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "2", "boundary": "census"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "2", "boundary": "cadastral"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "2", "boundary": "national_park"}},

		{[]int{0, 0, 0, 0, 0, 0, 1}, element.Tags{"name": "bmaritime", "boundary": "maritime"}},
		{[]int{0, 0, 0, 0, 0, 0, 1}, element.Tags{"name": "badmin", "boundary": "administrative"}},

		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "2", "name": "maritime2"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "3", "name": "maritime3"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "4", "name": "maritime4"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "5", "name": "maritime5"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "6", "name": "maritime6"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "7", "name": "maritime7"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "8", "name": "maritime8"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "X", "name": "maritimeX"}},
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "Y", "name": "maritimeY"}},

		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"name": "Helló World  öüóőúéáűí"}},

		// side effect test - __nil__ has a duble meaning ;   first check as a value and second as a NOT operator.
		{[]int{0, 0, 0, 0, 0, 0, 0}, element.Tags{"admin_level": "__nil__", "boundary": "__nil__"}},
		{[]int{0, 0, 0, 0, 0, 1, 0}, element.Tags{"admin_level": "__any__", "boundary": "__any__"}},

		{[]int{0, 0, 0, 0, 0, 1, 0}, element.Tags{"admin_level": "+", "boundary": "+"}},
		{[]int{0, 0, 0, 0, 0, 1, 0}, element.Tags{"admin_level": "-", "boundary": "-"}},
		{[]int{0, 0, 0, 0, 0, 1, 0}, element.Tags{"admin_level": "ナ", "boundary": "ナ"}},
	}

	// Query the number of test tables from the array
	testTableNumber := len(wayfilterTests[0].expected)
	fmt.Println("testTableNumber", testTableNumber)
	matched := make([]int, testTableNumber)

	elem := element.Way{}
	ls := configTestMapping.LineStringMatcher()

	var actual_match []Match

	for _, tt := range wayfilterTests {
		fmt.Println(tt)

		elem = element.Way{}
		ls = configTestMapping.LineStringMatcher()

		elem.Tags = tt.inputtags
		actual_match = ls.MatchWay(&elem)

		// fill matched with zero
		for i := 0; i <= (testTableNumber - 1); i++ {
			matched[i] = 0
		}

		for _, mt := range actual_match {
			nn, _ := strconv.Atoi(mt.Table.Name[1:2])
			matched[nn] = 1
		}

		for i := 0; i <= (testTableNumber - 1); i++ {
			if matched[i] != tt.expected[i] {
				t.Errorf(" TestAdvancedFiltering testcase[ elem.Tags:(%+v)  table:t%d   expected: %d,  actual: %d ] ", elem.Tags, i, tt.expected[i], matched[i])
				fmt.Println(i, " -- problem here!!!!")
			}

		}

	}

}
