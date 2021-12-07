# track
Time tracker for nerds

## Installation

### Binaries

Grab binaries for different OS from https://github.com/neonxp/track/releases

### Homebrew

```
brew install neonxp/tap/track
```

### With golang
```
go install github.com/neonxp/track@latest
```

## Usage

1. Add new activity:
```
track add Activity summary [#tag1 #tag2 @context1 @context2]
```
example:
```
~ track add Work on time tracker #tracker @home
Activity #1 added! Now you can start it.
```
2. Start activity:
```
track start ID [comment]
```
example:
```
~ track start 1 Writing documentation.
Started new span for activity "Work on time tracker".
```
3. List activities:
```
track ls [--all]
```
example:
```
~ track ls
Started activities:
1. Work on time tracker
        1 timespans
        19:17 5.12.2021 — now (10 minutes)
```
4. Complete work on activity. if activity id empty - stop all started activities:
```
track stop [ID]
```
example:
```
~ track stop
Stopped activity "Work on time tracker".
Last duration: 12 minutes.
All spent time: 12 minutes.
```
