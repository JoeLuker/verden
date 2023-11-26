export interface FormField {
    type: string;
    defaultValue: number | string;
    placeholder: string;
  }

export interface FormStructure {
    [key: string]: FormField;
  }

export interface FormData {
    [key: string]: number | string;
  }
