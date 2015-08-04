# `probe` design overview

The v0.1.0 will be focused on getting a basic use cases working end to end.
Following features have to be implemented
1. Store github access token with username
2. List assigned Pull requests for the user
3. Patch current branch with pull request
4. Remove patch and revert to original state
5. Show pull request information
6. Merge/Reject a pull request
7. Comment on pull request

In an attempt to adhere to "separation of concerns" philosophy, we'll try to
keep each component in a module

Command Parser  
==
- Entrypoint for the application
- supports following commands

Storage  
==
- expose API to abstract out getting/setting of storage data
- create a .probeconfig directory in `~/home directory`
- store a `config` object `~/home/.probeconfig/config` in json format
```
  {
    "userContext":{
      "github": {
        "<profile 1>": {
          "accessToken ": <profile1 access token>,
          "created": <unix TS for created at>
        },
        "profile 2": {
          "accessToken": <profile 2 access token>,
          "created": <unix TS for created at>
        }
      }
    },
    "projectContext": {
      "<path to project>": {
        "remotes": {
          "origin": {
            "url ": "git@github.com:myuser2/myproject.git
          },
          "github": {
            "url ": "git@github.com:myuser/myproject.git
          },
          "bitbucket": {
            "url ": "git@bb.com:myuser/myproject.git
          }
        }
        "remoteContext": "github"
      }
    }
  }
```
- prompt user for profile info if not present
- execute all commands within a profile context

Syncer
==
- expose API to synchronize content with the provider(github for now)
- support following calls
  - listRemotes()
  - setRemote(remoteName)
  - getAllRequests()
  - getRequest(requestId)
  - updateRequest(updateObj)
  - approveRequest(approveObj)
  - rejectRequest(rejectObj)

Actor
==
- expose API to execute commands in current directory
- support following calls
  - list-remotes
  - set-remote
  - viewall
  - view <id>
  - assigned
  - apply <id>
  - approve
  - reject
  - comment
