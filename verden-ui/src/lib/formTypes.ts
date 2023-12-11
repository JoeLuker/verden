export interface FormField {
  name: string;           // Added to match the JSON structure
  type: string;
  defaultValue: number;
  placeholder: string;
}

export interface Category {
  categoryName: string;   // Updated to match the JSON key
  fields: FormField[];    // Updated to lowercase 'f' to match the JSON key
  // other properties of Category
}

export interface FormStructure {
  categories: Category[]; // This already matches the JSON structure
}

export interface FormData {
  [key: string]: number | string; // This seems fine as is
}
