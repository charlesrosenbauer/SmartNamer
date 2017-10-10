# SmartNamer

A Simple Command Line Tool For Name Suggestions

---

Maybe you've heard the saying "There's only two hard things in computer science; cache invalidation and naming things." Well here we have a tool aimed at solving that second problem.



**How will/does it work?**

* Identifiers are detected in your code using regexes.
* A form of simple Semantic Folding is used to create sparse bit vectors for each identifier in your code, where similar ids (similar based on name and proximity in your code) will yield similar representations.
* Several tables consisting of (bit vector, identifier) pairs are searched for the best matches. Results are concatenated and displayed as potential names.
* A single-layer perceptron is used to take identifiers and predict names that better fit the surrounding code.


**How well does it work?**

Well it's blatantly imperfect, but it's clearly better than random. Currently when run on its own code, it makes a lot of predictions like "returnString","forString", and "returnFor". This is probably because those words show up a lot in the code. It could clearly use some work.
