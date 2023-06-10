# ___File Manager___ __(fm)__

__File Manager (fm)__ is a command-line tool for managing files and directories. It provides functionality such as copying, moving, renaming, and removing files and directories.

## __Installation__

To install and run File Manager, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/SharonAliyas5573/fm-cli.git

2. Change to the project directory:
    ```shell
    cd fm-cli

3. Build the project:
    ```shell
    make build

4. Install the binary to /bin with root privileges:
    ```shell
    sudo make install

5. File Manager is now installed and ready to use.

## __Usage__

File Manager supports the following commands:

- Copy files/directories:
    
      fm -cp [source] [destination]

- Move files/directories:

      fm -mv [source] [destination]
- Rename files:

      fm -rn [current_name] [new_name]

- Remove files/folders to trash:

      fm -rm [file/directory name]
### For more information on available commands and their options, run:
    fm -h


## __Examples__

1. Copy a file:

    ```shell
    fm -cp file.txt /path/to/destination

2. Move a directory:
    ```shell
    fm -mv directory /path/to/destination

3. Rename a file:
    ```shell
    fm -rn old_name.txt new_name.txt

4. Remove a file to the trash:
    ```shell
    fm -rm file.txt

## __FAQ__

### Q: Can I permanently delete a file without moving it to the trash?
A: Yes, you can use the -f flagwith the remove command. For _example_:

    fm -rm -f file.txt

## Troubleshooting

If you encounter any errors or issues while running File Manager, try the following:
- Ensure you have the necessary permissions to access the files and directories.
- Make sure the source and destination paths are correct and exist.
- Check your command syntax and arguments for any mistakes.

If you need further assistance, please create an issue on the project's GitHub repository.

### __Contributing__
__Contributions__ to File Manager are welcome! If you find a bug, have a suggestion, or want to contribute code, please submit a pull request on the project's GitHub repository.
License

This project is licensed under the __MIT License__.
