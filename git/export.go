package git

import (
	"context"
	"os"
	"path/filepath"
)

type Export struct {
	dir string
}

func (e *Export) Dir() string {
	return e.dir
}

func (e *Export) Clean() {
	if e.dir != "" {
		os.RemoveAll(e.dir)
	}
}

// Export creates a minimal clone of the repo, at the ref given.
func (r *Repo) Export(ctx context.Context, ref string) (*Export, error) {
	dir, err := r.workingClone(ctx, "")
	if err != nil {
		return nil, err
	}
	if err = checkout(ctx, dir, ref); err != nil {
		return nil, err
	}
	return &Export{dir}, nil
}

// ChangedFiles does a git diff listing changed files
func (c *Export) ChangedFiles(ctx context.Context, sinceRef string, paths []string) ([]string, error) {
	list, err := changed(ctx, c.Dir(), sinceRef, paths)
	if err == nil {
		for i, file := range list {
			list[i] = filepath.Join(c.Dir(), file)
		}
	}
	return list, err
}
