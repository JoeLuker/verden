import os

def save_directory_contents_to_file(path, output_file_path):
    with open(output_file_path, 'w', encoding='utf-8') as output_file:
        for root, dirs, files in os.walk(path):
            level = root.replace(path, '').count(os.sep)
            indent = ' ' * 4 * level
            output_file.write(f'{indent}{os.path.basename(root)}/\n')
            subindent = ' ' * 4 * (level + 1)
            for f in files:
                output_file.write(f'{subindent}{f}\n')
                file_path = os.path.join(root, f)
                with open(file_path, 'r', encoding='utf-8', errors='ignore') as file:
                    contents = file.read()
                    output_file.write(f'{subindent}Contents:\n{contents}\n\n')

# Replace './backend/simulation' with your desired directory and 'output.txt' with your desired output file path
save_directory_contents_to_file(r'./frontend/src/routes/simulation', 'output.txt')
