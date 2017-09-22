#SmartNamer

A Simple Command Line Tool For Name Suggestions

---

Maybe you've heard the saying "There's only two hard things in computer science; cache invalidation and naming things." Well here we have a tool aimed at solving that second problem.

**Note: This project is incomplete. Don't expect it to work yet.**



**How will/does it work?**

* Identifiers are detected in your code using regexes.
* A form of simple Semantic Folding is used to create sparse bit vectors for each identifier in your code, where similar ids (similar based on name and proximity in your code) will yield similar representations.
* Several tables consisting of (bit vector, identifier) pairs are searched for the best matches. Results are concatenated and displayed as potential names.

**Can I Help?**

If you know Haskell, sure! Contact me if you want info on what you can help with.
