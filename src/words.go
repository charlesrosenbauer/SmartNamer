package main










var words []string = []string{ "map", "filter", "reduce", "fold", "scan",
  "bind", "zip", "unzip", "define", "either", "maybe", "functor", "applicative",
  "monad", "promise", "future", "apply", "call", "base", "identity", "identify",
  "compose", "query", "search", "watch", "show", "able", "unable", "usable",
  "display", "draw", "compile", "create", "form", "delta", "compress",
  "compressing", "compressor", "for", "while", "name", "identifier", "fail",
  "failure", "test", "error", "signal", "exception", "except", "accept",
  "catch", "catcher", "throw", "thrower", "pass", "succeed", "socket", "lead",
  "leave", "leaving", "quit", "mail", "email", "format", "print", "fetch",
  "goal", "mark", "anneal", "annealing", "annealer", "rate", "rating", "pack",
  "package", "repo", "repository", "string", "bool", "boolean", "int", "integer",
  "float", "floating", "point", "double", "set", "hashmap", "matrix", "learn",
  "learning", "tensor", "deep", "order", "fizz", "buzz", "beep", "sound",
  "play", "stop", "clock", "pause", "use", "using", "done", "complete",
  "completed", "resume", "trigger", "represent", "insert", "remove", "trade",
  "exchange", "change", "report", "repair", "mix", "id", "bin", "binary", "hex",
  "handle", "handler", "smart", "get", "set", "alloc", "allocate", "pointer",
  "array", "hash", "len", "length", "size", "import", "importer", "struct",
  "specific", "specify", "spec", "factor", "find", "replace", "attach",
  "connect", "standard", "std", "my", "your", "me", "you", "all", "none", "yet",
  "swap", "quick", "rapid", "insertion", "bubble", "bogo", "sort", "view",
  "user", "users", "if", "then", "else", "gate", "race", "condition", "compare",
  "comp", "load", "store", "cache", "recog", "recognition", "recognize",
  "recognizer", "match", "matcher", "pattern", "dictionary", "dict", "diction",
  "table", "array", "list", "link", "graph", "tree", "dag", "ast", "syntax",
  "add", "subtract", "sub", "multiply", "mul", "divide", "div", "sum",
  "product", "quotient", "remainder", "remain", "modulus", "mod", "addition",
  "additional", "vector", "matrix", "cross", "dot", "volume", "perimeter",
  "area", "height", "width", "depth", "associate", "associative", "commute",
  "commutative", "algebra", "calculus", "category", "theory", "info",
  "information", "entropy", "entropic", "trig", "sine", "cosine", "tangent",
  "arc", "sin", "cos", "tan", "count", "accum", "accumulate", "arc", "arch",
  "parabola", "parabolic", "auto", "struct", "break", "switch", "case", "enum",
  "register", "type", "def", "definition", "char", "character", "extern",
  "external", "return", "union", "cont", "continue", "sign", "signed", "void",
  "nil", "null", "do", "dont", "static", "default", "goto", "size", "volatile",
  "const", "constant", "short", "long", "unsigned", "account", "balance",
  "setting", "settings", "pre", "post", "before", "after", "ignore", "pipe",
  "pipeline", "latency", "latent", "bandwidth", "channel", "routine", "co", "go",
  "rotate", "rot", "nest", "nested", "line", "root", "sqrt", "cbrt", "log",
  "radical", "rad", "ratio", "rational", "real", "number", "num", "numerical",
  "precision", "precise", "prec", "ration", "exp", "exponent", "exponential",
  "complex", "compound", "poly", "polygon", "gon", "hyper", "super", "uber",
  "extra", "quartic", "quart", "schedule", "scheduler", "process", "processor",
  "thread", "core", "task", "state", "fiber", "recursive", "rec", "recurse",
  "group", "queue", "stack", "fifo", "lifo", "fofi", "lofi", "aspect", "compute",
  "calculate", "calc", "calculus", "run", "runtime", "emit", "omit", "emitter",
  "omitter", "factory", "data", "regex", "regexpr", "reg", "regular", "regulate",
  "regularize", "normal", "normalize", "legal", "bit", "byte", "effect",
  "affect", "cons", "car", "cdr", "cns", "snoc", "tacit", "reverse", "rev",
  "type", "var", "variable", "fn", "fun", "func", "function", "dep", "depend",
  "dependency", "dependent", "arg", "argument", "arguments", "param", "par",
  "parameter", "macro", "method", "template", "temp", "temporary", "obj",
  "object", "class", "alias", "job", "app", "application", "serve", "service",
  "server", "client", "cover", "covering", "covered", "status", "stat",
  "statistic", "intern", "internal", "hire", "fire", "master", "slave",
  "fluffykins", "observer", "observe", "misc", "username", "password", "pass",
  "word", "name", "proto", "prototype", "abstract", "crypt", "crypto", "encrypt",
  "decrypt", "encode", "decode", "code", "encryption", "decryption", "cypher",
  "encypher", "decypher", "rsa", "aes", "sha", "sha1", "sha2", "sha3", "async",
  "sync", "para", "parallel", "concurrent", "current", "pool", "lazy", "strict",
  "eager", "facade", "proxy", "composite", "decorate", "decorate", "bridge",
  "adapt", "adapter", "wrap", "wrapper", "trans", "translate", "translator",
  "iter", "iterate", "iterator", "chain", "active", "inactive", "block",
  "blockchain", "join", "lock", "free", "dealloc", "deallocate", "malloc",
  "calloc", "message", "send", "receive", "receiver", "tranceiver", "record",
  "recorder", "listen", "explode", "assert", "panic", "defer", "deferred",
  "mediate", "mediator", "mutex", "semaphore", "flyweight", "proximal", "distal",
  "local", "global", "nonlocal", "and", "xor", "not", "or", "non", "with",
  "without", "in", "out", "ex", "shift", "logic", "logical", "point", "to",
  "from", "push", "pop", "enqueue", "deque", "dequeue", "inter", "intersect",
  "section", "diff", "difference", "different", "differential", "integral",
  "integ", "integrate", "deviate", "deviant", "deviation", "lambda", "expr",
  "expression", "gauss", "gaussian", "euler", "eulerian", "newton", "newtonian",
  "hilbert", "peano", "serpinski", "turing", "cantor", "euclid", "euclidean",
  "noneuclidean", "linear", "nonlinear", "tietze", "homeo", "morphism", "homo",
  "hetero", "multi", "mono", "biject", "bijection", "inject", "injection",
  "compact", "hausdorff", "fractal", "fraction", "knot", "gosper", "glider",
  "gun", "conway", "game", "life", "koch", "madel", "mandelbrot", "moore",
  "murray", "osgood", "os", "system", "logo", "logos", "tag", "tagged", "tagger",
  "pos", "posit", "position", "loc", "locate", "location", "design", "designer",
  "cost", "equate", "equation", "finance", "money", "percent", "money", "euro",
  "dollar", "yen", "coin", "bitcoin", "ether", "force", "magnitude", "abs",
  "absolute", "box", "talk", "dialog", "dialogue", "dial", "look", "hide",
  "save", "guard", "kill", "revive", "heal", "fix", "heavy", "light", "big",
  "small", "protocol", "internet", "network", "vision", "http", "tcp", "udp",
  "json", "xml", "yaml", "js", "javascript", "script", "lua", "java", "golang",
  "go", "rust", "python", "bzo", "haskell", "lisp", "clojure", "jai", "scala",
  "node", "c", "cpp", "swift", "markdown", "github", "git", "hub", "ignore",
  "license", "readme", "gpl", "bsd", "mit", "apache", "runtime", "gc", "garbage",
  "collection", "collector", "manager", "library", "language", "opencl", "cl",
  "opengl", "gl", "openal", "al", "opencv", "cv", "vulkan", "vk", "metal", "mt",
  "mantle", "directx", "dx", "intel", "amd", "amd64", "x86", "x64", "i386",
  "i686", "sdl", "qt", "byte", "hword", "dword", "measure", "detect", "prefix",
  "token", "symbol", "adopt", "base", "top", "bottom", "side", "front", "back",
  "backspace", "space", "less", "great", "greater", "equal", "inequal",
  "invalid", "express", "resolve", "solve", "eval", "evaluate", "than", "this",
  "near", "unknown", "far", "memory", "unusual", "common", "unit", "grammar",
  "lexicon", "parser", "parse", "lexer", "lex", "text", "semantic", "file",
  "operate", "operating", "loan" }