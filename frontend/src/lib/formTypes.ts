export interface FormField {
    type: string;
    defaultValue: number | string;
    placeholder: string;
  }

export interface FormStructure {
  categories: Category[];
}
  

export interface FormData {
    [key: string]: number | string;
  }

export interface Category {
    Fields: FormField[];
    // other properties of Category
  }
