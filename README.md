# `glsnip` - CLI tool for GitLab Snippets

## Features

- [x] Create snippets
- [x] Delete snippets
- [ ] Update snippets
- [ ] List snippets

## Installation

```
go install github.com/blufor/glsnip
```

Make sure you have these ENV variables set-up:
- `GITLAB_URL`
- `GITLAB_TOKEN`

## Usage

The command displays all the parameters:

```
> glsnip -h
GitLab Snippet CLI util

Usage:
  glsnip [flags]

Flags:
  -D, --delete uint          Delete snippet by its ID
  -d, --description string   description of the snippet
  -h, --help                 help for glsnip
  -i, --internal             internal visibility
  -p, --public               public visibility
  -t, --title string         title for the snippet
      --token string         GitLab authentication token (use GITLAB_TOKEN in ENV)
  -U, --update uint          Update snippet by its ID
      --url string           GitLab URL (use GITLAB_URL in ENV)
```

### Create snippet

#### From `STDIN`

```
uname -a | glsnip -t 
```

#### From file

You can upload a single file:
```
glsnip -t test /proc/cpuinfo
```

or multiple files:

```
glsnip -t test /proc/cpuinfo /proc/vmstat
```

### Delete snippet

You can delete your snippet by its ID:

```
glsnip -D <ID>
```
