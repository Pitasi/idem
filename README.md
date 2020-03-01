# Idem

A simple recommend engine for anything - based on *likes*, Facebook-style.

This basic implementation came out while I was having a shower. There's plenty
of room for improvements - any suggestion is welcome, so please continue
reading.


## Rationale

The two entities are simple: users and *things*.

An user can like any number of things.

**Case 1:**
Based on the things the other users you want to suggest to a certain user
something it may like too.

**Case 2:**
Based on a thing, you want to list other things that are similar and that the
users may likes too.


## Todo

- [x] Case 1
- [ ] Case 2
- [ ] Reduce used memory by storing only half of the graph (instead of storing
  A->B and B->A we should exploit the symmetry of the problem to just store
  A-B).
- [ ] Build a simple webserver with a small set of HTTP endpoints (e.g. GET
  /suggestions/{user}, POST /like/{user}/{thing})
- [ ] Make some serialization/deserialization available for persistence. Still
  not sure if picking a dabatase, use something embedded like Bolt, or just be
  agnostic.


## Note

I don't need this to be based on some NN or be super-fast. I plan to use this
on a small set of available things and to update the cache of suggestions every
once in a while.
