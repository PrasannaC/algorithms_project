package Set

/*
	Structure of the set stored in disjoiint sets data structure
*/
type Set struct {
	Data   interface{}
	Parent *Set
	/* will be only set for the root, as ever element
	does not need to have this set in union-find
	*/
	Rank int
	Size int
}
