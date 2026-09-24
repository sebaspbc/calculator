import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import Calculator from "./Calculator";
import * as api from "../services/api";

describe("Calculator", () => {
  it("performs addition end to end via the API", async () => {
    vi.spyOn(api, "calculate").mockResolvedValue(8);
    render(<Calculator />);

    fireEvent.click(screen.getByText("5"));
    fireEvent.click(screen.getByText("+"));
    fireEvent.click(screen.getByText("3"));
    fireEvent.click(screen.getByText("="));

    await waitFor(() => {
      expect(screen.getByTestId("display")).toHaveTextContent("8");
    });
  });

  it("shows an error message when the API rejects the operation", async () => {
    vi.spyOn(api, "calculate").mockRejectedValue(new Error("division by zero is not allowed"));
    render(<Calculator />);

    fireEvent.click(screen.getByText("1"));
    fireEvent.click(screen.getByText("÷"));
    fireEvent.click(screen.getByText("0"));
    fireEvent.click(screen.getByText("="));

    await waitFor(() => {
      expect(screen.getByText("division by zero is not allowed")).toBeInTheDocument();
    });
  });

  it("adds an entry to history after a successful operation", async () => {
  vi.spyOn(api, "calculate").mockResolvedValue(8);
  render(<Calculator />);

  fireEvent.click(screen.getByText("5"));
  fireEvent.click(screen.getByText("+"));
  fireEvent.click(screen.getByText("3"));
  fireEvent.click(screen.getByText("="));

  await waitFor(() => {
    expect(screen.getByText("5 + 3")).toBeInTheDocument();
    expect(screen.getByText("= 8")).toBeInTheDocument();
  });
});
});