import type { CalculateRequest, CalculateResponse, ApiErrorResponse } from "../types/calculator.types";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

export async function calculate(payload: CalculateRequest): Promise<number> {
  const response = await fetch(`${API_URL}/api/v1/calculate`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  const data = await response.json();

  if (!response.ok) {
    const errorData = data as ApiErrorResponse;
    throw new Error(errorData.error || "Unexpected error occurred");
  }

  return (data as CalculateResponse).result;
}