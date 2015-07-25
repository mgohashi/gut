package main

import (
    "errors"
    "path"
    "strings"
)

// Query the gut repo for the initial commit to the repo. We use this to determine if two gut repos are compatibile.
// http://stackoverflow.com/questions/1006775/how-to-reference-the-initial-commit
func GetTailHash(ctx *SyncContext) (string, error) {
    exists, err := ctx.PathExists(path.Join(ctx.AbsSyncPath(), ".gut"))
    if err != nil { ctx.Logger().Bail(err) }
    if exists {
        output, err := ctx.GutOutput("rev-list", "--max-parents=0", "HEAD")
        if err != nil { ctx.Logger().Bail(err) }
        return strings.TrimSpace(output), nil
    }
    return "", nil
}

func GutDaemon(ctx *SyncContext, tailHash string, bindPort int) (err error) {
    return errors.New("Not implemented")
}

func GutInit(ctx *SyncContext) (err error) {
    return errors.New("Not implemented")
}

func GutSetupOrigin(ctx *SyncContext, tailHash string, connectPort int) (err error) {
    return errors.New("Not implemented")
}

func GutPull(ctx *SyncContext) (err error) {
    return errors.New("Not implemented")
}

func GutCommit(ctx *SyncContext, prefix string, updateUntracked bool) (changed bool, err error) {
    return false, errors.New("Not implemented")
}

func GutEnsureInitialCommit(ctx *SyncContext) (err error) {
    return errors.New("Not implemented")
}
