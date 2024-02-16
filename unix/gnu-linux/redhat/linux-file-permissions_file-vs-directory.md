# Linux File Permissions (_FILE_ vs _DIRECTORY_)

## What do Linux File Permissions Actually Do?

I've talked about how to view file permissions, who they apply to, and how to read what permissions are enabled or disabled. But what do these permissions actually do in practice?

### Read (r)

Read permission is used to access the file's contents. You can use a tool like cat or less on the file to display the file contents. You could also use a text editor like Vi or view on the file to display the contents of the file. Read permission is required to make copies of a file, because you need to access the file's contents to make a duplicate of it.

### Write (w)

Write permission allows you to modify or change the contents of a file. Write permission also allows you to use the redirect or append operators in the shell (> or >>) to change the contents of a file. Without write permission, changes to the file's contents are not permitted.

### Execute (x)

Execute permission allows you to execute the contents of a file. Typically, executables would be things like commands or compiled binary applications. However, execute permission also allows someone to run Bash shell scripts, Python programs, and a variety of interpreted languages.

There are other ways to execute the contents of a file without execute permission. For example, you could use an interpreter that has execute permission to read a file with instructions for the interpreter to execute. An example would be invoking a Bash shell script:

```sh
$ bash script.sh
```

The executable being run is bash. The script.sh file is read by the Bash interpreter, and its commands are executed. The content in this article is general purpose, but in Linux, there are often additional ways to accomplish tasks.

## How do Directory Permissions Work?

Directory file types are indicated with d. Conceptually, permissions operate the same way, but directories interpret these operations differently.

### Read (r)

Like regular files, this permission allows you to read the contents of the directory. However, that means that you can view the contents (or files) stored within the directory. This permission is required to have things like the ls command work.

### Write (w)

As with regular files, this allows someone to modify the contents of the directory. When you are changing the contents of the directory, you are either adding files to the directory or removing files from the directory. As such, you must have write permission on a directory to move (mv) or remove (rm) files from it. You also need write permission to create new files (using touch or a file-redirect operator) or copy (cp) files into the directory.

### Execute (x)

This permission is very different on directories compared to files. Essentially, you can think of it as providing access to the directory. Having execute permission on a directory authorizes you to look at extended information on files in the directory (using ls -l, for instance) but also allows you to change your working directory (using cd) or pass through this directory on your way to a subdirectory underneath.

Lacking execute permission on a directory can limit the other permissions in interesting ways. For example, how can you add a new file to a directory (by leveraging the write permission) if you can't access the directory's metadata to store the information for a new, additional file? You cannot. It is for this reason that directory-type files generally offer execute permission to one or more of the user owner, group owner, or others.

---

_REFERENCELINKS_:

- From: <https://www.redhat.com/sysadmin/linux-file-permissions-explained>.
