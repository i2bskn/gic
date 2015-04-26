gic (GitHub Issue Creator)
==========================

Create GitHub issue from template.

Installation
------------

With go lang:

```
go get github.com/i2bskn/gic
```

Settings
--------

gic is required [Personal Access Token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) of GitHub.  
Personal Access Token must be able to get in the `git-config` command.

```
[github]
  token = <personal access token>
```

Usage
-----

Initialize gic in project.  
Create `.gic` directory to project root.  
Not initialize if already been initialized.

```
cd /path/to/project
gic init
```

Create template.  
Templates is stored in `.gic/templates`.

```
gic edit <template name>
```

Example:

```
# Title

## Files

{{.Execute "ls -l"}}

This template was written by {{.Env.EDITOR}}.
```

Create GitHub Issue.

```
gic list # => Display a list of templates.
gic preview <template name> # => Display a preview of template.
gic apply <template name> # => Create issue to current project.
```

