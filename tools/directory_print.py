import os

def save_directory_contents_to_file(path, output_file_path):
    def is_excluded(name):
        """Determine if a file or directory should be excluded."""
        excluded_prefixes = ['tmp', '.', 'go.', 'node_', 'lib', 'yarn']
        return any(name.startswith(prefix) for prefix in excluded_prefixes)

    with open(output_file_path, 'w', encoding='utf-8') as output_file:
        for root, dirs, files in os.walk(path):
            # Exclude directories that match the exclusion criteria
            dirs[:] = [d for d in dirs if not is_excluded(d)]

            # Exclude files that match the exclusion criteria
            files = [f for f in files if not is_excluded(f)]

            # Write the directory name to the output file
            level = root.replace(path, '').count(os.sep)
            indent = ' ' * 4 * level
            output_file.write(f'{indent}{os.path.basename(root)}/\n')

            # Write the names and contents of the files to the output file
            subindent = ' ' * 4 * (level + 1)
            for f in files:
                output_file.write(f'{subindent}{f}\n')
                file_path = os.path.join(root, f)
                with open(file_path, 'r', encoding='utf-8', errors='ignore') as file:
                    contents = file.read()
                    output_file.write(f'{subindent}Contents:\n{contents}\n\n')

save_directory_contents_to_file(r'./frontend/src/routes/Simulation', './tools/output.txt')

