export type Operation =
  | "add"
  | "subtract"
  | "multiply"
  | "divide"
  | "power"
  | "sqrt"
  | "percentage";

export interface CalculateRequest {
  operation: Operation;
  a: number;
  b?: number;
}

export interface CalculateResponse {
  result: number;
}

export interface ApiErrorResponse {
  error: string;
}

export interface HistoryEntry {
  id: string;
  expression: string;
  result: number;
}