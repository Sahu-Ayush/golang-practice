✅ Final Structure: go-practice Repo

go-practice/
├── go.mod                          # module github.com/ayushsahu/go-practice
├── README.md                       # Full guide, links, setup
├── tools/
│   └── check_module_name.go        # Validates module name
├── scripts/
│   └── pre-push-check.sh           # Optional Git hook for module validation
├── .git/hooks/
│   └── pre-push                    # Symlink to scripts/pre-push-check.sh (optional)
├── phase1_basics/
│   ├── 01_hello.go
│   ├── 02_variables.go
│   ├── ...
│   └── 07_practice_exercises/
│       ├── ...
├── playground_links.md             # All Go Playground pre-filled links
└── phase_summary.pdf               # (Optional) Printable PDF
