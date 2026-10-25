# term-in-algebra -- a linear algebra workbench for the terminal

Your helper for understanding (some) linear algebra operations.
Performs the operation, computes the result and displays the calculation steps.

## Demo

## Install & Run

## Usage

## Operations

The definitions, worked examples and numerical pitfalls behind every operation
are in [docs/math-reference.md](docs/math-reference.md).

## Roadmap

Phases match Parts 1 and 2 of the math reference. Within each phase, items are
in build order: later items reuse earlier ones.

### Phase 1 — entrywise and product basics

- [ ] matrix input/parsing and terminal display (needed by everything else)
- [ ] identity
- [ ] add / subtract
- [ ] scalar multiply
- [ ] transpose
- [ ] trace
- [ ] elementwise ops (Hadamard product, etc.)
- [ ] matrix-vector multiply
- [ ] norms (vector 1/2/∞; matrix Frobenius, 1, ∞)

### Phase 2 — elimination and everything built on it

- [ ] matrix multiply
- [ ] row operations as user commands (swap, scale, `R_i ← R_i + c·R_j`)
- [ ] RREF / Gauss-Jordan, showing every row operation
- [ ] rank (pivot count from RREF)
- [ ] determinant (by elimination; cofactor expansion shown step by step only for n ≤ 3)
- [ ] inverse (reduce `[A | I]` to `[I | A⁻¹]`)
- [ ] solve Ax=b, reporting which case applies: unique, infinitely many (general solution with named free variables), or none
- [ ] LU with partial pivoting

### Output

- [ ] Unicode/ASCII matrices in the terminal
- [ ] ANSI 256/truecolor shaded cells by value; sparklines/bars for vectors

## Design decisions

- **Exact rational arithmetic by default** (Go's `math/big.Rat`). Steps show
  `R2 ← R2 - 3/2·R1`, not `R2 ← R2 - 1.4999999·R1`, and rank, RREF and
  determinants don't depend on a floating-point tolerance. A float mode may
  come later; norms involving square roots need a float (or `√n`) display path.
- **Steps are data, not print statements.** Every operation returns its result
  plus a list of steps (what was done, the matrix after it, a short
  explanation). A separate renderer turns that into terminal output, so the
  maths stays testable and the display can change independently.
- **No broadcasting.** Dimension mismatches are errors whose message names both
  shapes.
- **Hadamard is never called "multiplication"** in the interface, to keep it
  distinct from the matrix product.

## Non-goals (for now)

- Eigenvalues/eigenvectors, SVD, the spectral 2-norm, condition numbers
- QR, Cholesky and other factorizations beyond LU
- Least squares and over/underdetermined systems beyond the three `Ax = b` cases
- Complex numbers, sparse matrices, performance work (BLAS/LAPACK, GPU)

See Part 4 of the math reference for the full list of v2 candidates.

## Development

Planned Go layout:

```text
cmd/tia/          entry point (REPL + one-shot CLI)
internal/matrix/  matrix type and arithmetic (big.Rat)
internal/ops/     operations that return a result plus steps
internal/parse/   matrix and command input
internal/render/  turns steps into terminal output
```

Testing: every worked example in the math reference becomes a table-driven
test; identities such as `(AB)ᵀ = BᵀAᵀ`, `A·A⁻¹ = I` and
`det(AB) = det(A)·det(B)` are checked on random rational matrices; rendered
output is checked with golden files.
