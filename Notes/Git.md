0. git basics
    - status
        - untracked: new file added
        - modified: edited existing file
        - staged: changes that "git add"-ed
        - committed
        - pushed

1. git init
    - Initializes a new Git repository in the current directory.
    - use -b (git init -b newName) to specify name of initial branch

2. git clone
    - clone existing repo
    - clone
        - git clone https://github.com/username/repo.git
        - git clone local/dir/folder
    - clone branch
        - git clone -b feature-branch from.git
    - clone to own repo
        - git clone from.git my/dir
    - clone with depth limit ( number of commit history)
        - git clone --depth 1 from.git

3. git status
    - show information of local repo, include
        - on which branch
        - staged changes ( after git add )
        - not staged changes + untracked file ( after changing document)
        - unpushed changes

4. git add
    - stage changes for next commit
    - stage 1 file
        - git add file1.txt
    - stage multiple file
        - git add file1.txt file2.txt
    - stage all (new, modified, deleted)
        - git add . ( from git v2 onwards)
        - git add -A
    - stage (modified, deleted)
        - git add -u
    - stage (new, modified)
        - git add --ignore-removal .
    - stage with wild cards
        - git add *.txt

5. git commit
    - commit all changes
        - git commit -m "comment"
    - commit specific file
        - git commit file1.txt file2.txt -m "comment"
    - commit with multi line comment
        - git commit -m "first line" -m "second line"
    - update last commit message/ add new changes
        - git commit --amend -m "updated"
    - stage and commit ( except new file)
        - git commit -am "message"
    - commit with template
        - git commit -t path/to/template.txt
    - verbose mode
        - git commit -v -m "comment"

6. git push
    - push to default
        - git push
    - push for new branch ( first time )
        - git push -u origin branch
        - "u" to set upstream
    - push specific branch 
        - git push origin branch1 ( push from local branch1 to remote branch1 )
    - push all local branch
        - git push --all
    - force change ( overwrite others' change)
        - git push --force
            - example: git checkout oldertag, git reset --hard older tag, git push --force
        - use when need to rewtie history ( like rebase, amend, squashing commit, reset)
            - git push --force-with-lease
                -  when commit after rebased branch ( to make sure no one commit after last check) 
            - git push --force-if-includes
                - when commit after rebased branch, but others have committed based rebased commit
                - so we fetch/pull/merge in the change and use this comment to makesure others work preserved
    - push tags
        - git push origin v1.0

7. git branch
    - list out branches
        - git branch
    - create new branch
        - git branch <branch-name>
    - list all remote branch
        - git branch -r
    - delete branch
        - git branch -d

8. git checkout
    - check out branch
        - git checkout <branch-name>
    - create and new branch and checkout
        - git checkout -b <new-branch>
    - checkout specific commit
        - git checkout <commit-hash> ( can get from git reflog)
    - checkout certain file
        - git checkout <commit-hash> -- <file-path>
    - checkout file from last commit
        - git checkout -- <file-path>

9. git merge
10. git rebase
11. git log
11. git diff
12. git reset
13. git revert
14. git stash
    - git stash pop

15. git remote
    - git remote add

16. git tag
    - add tag on specific commit
        - git tag <tagname>
    - annotated tag ( with message )
        - git tag -a <tagname> -m " message"
    - list out all tag
        - git tag

17. git cherry-pick
18. git config
19. git <command> --help
20. git squash commit
21. git reflog
22. git switch
    - git switch <branch-name>
    - switch ebtween branch
    