import { useState } from "react";
import { calculate } from "../services/api";
import type { Operation, HistoryEntry } from "../types/calculator.types";
import "./Calculator.css";

const BINARY_OPS: Record<string, Operation> = {
  "+": "add",
  "-": "subtract",
  "×": "multiply",
  "÷": "divide",
  "^": "power",
};

export default function Calculator() {
  const [display, setDisplay] = useState("0");
  const [storedValue, setStoredValue] = useState<number | null>(null);
  const [pendingOp, setPendingOp] = useState<Operation | null>(null);
  const [overwrite, setOverwrite] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [history, setHistory] = useState<HistoryEntry[]>([]);

  const clearError = () => setError(null);

  const handleDigit = (digit: string) => {
    clearError();
    if (overwrite) {
      setDisplay(digit === "." ? "0." : digit);
      setOverwrite(false);
      return;
    }
    if (digit === "." && display.includes(".")) return;
    setDisplay((prev) => (prev === "0" && digit !== "." ? digit : prev + digit));
  };

  const handleClear = () => {
    setDisplay("0");
    setStoredValue(null);
    setPendingOp(null);
    setOverwrite(true);
    clearError();
  };

  const resolve = async (a: number, b: number, op: Operation, nextOp?: Operation) => {
    setLoading(true);
    try {
      const result = await calculate({ operation: op, a, b });
      setDisplay(String(result));
      setStoredValue(nextOp ? result : null);
      setPendingOp(nextOp ?? null);
      setOverwrite(true);

      const symbolMap: Partial<Record<Operation, string>> = {
        add: "+", subtract: "−", multiply: "×", divide: "÷", power: "^",
      };
      const symbol = symbolMap[op] ?? op;
      setHistory((prev) => [
        { id: crypto.randomUUID(), expression: `${a} ${symbol} ${b}`, result },
        ...prev,
      ].slice(0, 10)); 
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unexpected error");
      setOverwrite(true);
    } finally {
      setLoading(false);
    }
  };

  const handleBinaryOperator = (symbol: string) => {
    clearError();
    const op = BINARY_OPS[symbol];
    const current = parseFloat(display);

    if (storedValue !== null && pendingOp && !overwrite) {
      void resolve(storedValue, current, pendingOp, op);
    } else {
      setStoredValue(current);
      setPendingOp(op);
      setOverwrite(true);
    }
  };

  const handleEquals = () => {
    if (storedValue === null || pendingOp === null) return;
    void resolve(storedValue, parseFloat(display), pendingOp);
  };

  const handleUnary = async (op: Operation) => {
    clearError();
    const current = parseFloat(display);
    setLoading(true);
    try {
      const result = await calculate({ operation: op, a: current });
      setDisplay(String(result));
      setOverwrite(true);

      const label = op === "sqrt" ? `√${current}` : `${current}%`;
      setHistory((prev) => [
        { id: crypto.randomUUID(), expression: label, result },
        ...prev,
      ].slice(0, 10));
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unexpected error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="calculator-wrapper">
      <div className="calculator" role="group" aria-label="Calculator">
        <div className="calculator__display" data-testid="display">
          {display}
        </div>
        {error && (
          <div className="calculator__error-banner" role="alert">
            {error}
          </div>
        )}
        <div className="calculator__grid">
          <button onClick={handleClear}>AC</button>
          <button onClick={() => handleUnary("sqrt")}>√</button>
          <button onClick={() => handleUnary("percentage")}>%</button>
          <button onClick={() => handleBinaryOperator("÷")}>÷</button>

          {["7", "8", "9"].map((d) => (
            <button key={d} onClick={() => handleDigit(d)}>{d}</button>
          ))}
          <button onClick={() => handleBinaryOperator("×")}>×</button>

          {["4", "5", "6"].map((d) => (
            <button key={d} onClick={() => handleDigit(d)}>{d}</button>
          ))}
          <button onClick={() => handleBinaryOperator("-")}>−</button>

          {["1", "2", "3"].map((d) => (
            <button key={d} onClick={() => handleDigit(d)}>{d}</button>
          ))}
          <button onClick={() => handleBinaryOperator("+")}>+</button>

          <button onClick={() => handleDigit("0")}>0</button>
          <button onClick={() => handleDigit(".")}>.</button>
          <button onClick={() => handleBinaryOperator("^")}>^</button>
          <button onClick={handleEquals} disabled={loading} className="calculator__equals">
            {loading ? "…" : "="}
          </button>
        </div>
      </div>

      {history.length > 0 && (
        <aside className="history" aria-label="Operation history">
          <div className="history__header">
            <span>History</span>
            <button className="history__clear" onClick={() => setHistory([])}>Clear</button>
          </div>
          <ul className="history__list">
            {history.map((entry) => (
              <li key={entry.id}>
                <span className="history__expr">{entry.expression}</span>
                <span className="history__result">= {entry.result}</span>
              </li>
            ))}
          </ul>
        </aside>
      )}
    </div>
  );
}