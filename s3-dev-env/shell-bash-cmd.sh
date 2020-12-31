# Print Working Directory
pwd

# List
ls

# List - list all information
ls -la
# d = directory
# rwxrwxrwx = owner, group, world (read, write, execute)

# Change Directory
cd
cd ../ # Move to parent directory

# Create a file
touch {file_name1} ...

# Edit file
nano {file_name}

# Print file contents
cat {file_name}

# Make Directory
mkdir {dir_name1} {dir_name2} ...

# Clear output
clear

# Change Permissions
chmod

# Show Environment Variables
env

# Remove file
rm {file_name}

# Remove recursively with force
rm -rf {file_name} {folder_name}

# .bash_profile & .bashrc
# .bash_profile is executed for login shells
# .bashrc is executed for interactive non-login shells

# Search for sth
grep
cat {file_name} | grep {search}
ls | grep -i {search}
