# Tanda - A package versioning command line tool

Have you ever use `npm version` in your JavaScript project? I love to that command line. I hope there is a similar one for flutter project.

So, I build this tool. 

`tanda` will changes the version your Flutter project by changing data in `pubspec.yaml` and also create git commit & tag for the version.

## How to install

### Snapcraft on Ubuntu (other linux that has snapcraft)

```console
$ sudo snap install tanda
```

## How to use

1. Your flutter project must have `pubspec.yaml`.
2. Run command `tanda`, it will show:
    
    ```bash
    $ tanda
    [12:05:05] INFO   tanda version 0.1.0

    [12:05:05] INFO   NAME:     isave_app
    [12:05:05] INFO   VERSION:  4.1.1+72
    [12:05:05] INFO   TYPE:     Flutter
    ```
3. As you can see, first line it showing `tanda` tool version. Next three lines, it showing detail of your Flutter Project.
4. Use command `tanda patch` to create new patched version.
   
    ```bash
    $ tanda patch
    [12:07:39] INFO   tanda version 0.1.0
    [12:07:39] INFO   Patch version bump for Flutter
    [12:07:39] INFO   Bump the `isave_app` from 4.1.1+72 to 4.1.2+73
    ```
5. As you can see, first line is `tanda` cli version.
6. Next line, it tells you that `tanda` tool is updating the version for Flutter. Why it tell you Flutter? Because, this same tool is planned so that can be use on NPM (Javascript) package too. (But not working yet)
7. Then in the last line, it tell you it update from version `4.1.1+72` to `4.1.2+73`.
8. It update the patch part on version name from `4.1.1` to `4.1.2` & it also on version code increased from `72` to `73`.
9. Plus, if you see the `git log` you will see new commit & tag is created for the updated version.
10. So, `tanda` will make your job updating Flutter project version easier.

## Documentation

- `tanda` - Show current package
- `tanda major` - Bump major version
- `tanda minor` - Bump minor version
- `tanda patch` - Bump patch version
