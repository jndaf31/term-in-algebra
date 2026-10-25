# Linear algebra operations — reference for `term-in-algebra`

A refresher for someone who has studied linear algebra before and needs the
definitions, the notation, and a worked example back in working memory.

---

## 0. Notation and conventions

- A matrix is written `A`, its entries `a_ij` — row `i`, column `j`, both
  starting at 1 (or 0 in code; pick one and stay consistent).
- Dimensions are written `m × n` (rows × columns). A vector is an `n × 1`
  column matrix unless stated otherwise.
- `Aᵀ` transpose, `A⁻¹` inverse, `I_n` identity, `0` zero matrix.
- `R_i` means row `i`; `a_ij` means the entry at row `i`, column `j`.
- `Σ` means sum; `∏` means product; `|x|` absolute value; `det(A)` or `|A|`
  determinant.
- `e_i` is the `i`-th standard basis vector (a 1 in position `i`, zeros
  elsewhere).
- Complexity is given in flops (floating-point operations) for an `n × n`
  matrix unless stated otherwise.

Throughout, examples use small integers so you can verify results by hand.

---

# Part 1 — Phase 1 operations

## 1. Addition and subtraction

**Notation:** `A + B`, `A - B`

**Definition:** Entrywise, and only defined when the dimensions match exactly.
For `A`, `B` both `m × n`:

```text
(A + B)_ij = a_ij + b_ij
(A - B)_ij = a_ij - b_ij
```

Subtraction is addition of the negation: `A - B = A + (-1)·B`.

**Example:**

```text
[1 2]   [5 6]   [6  8]
[3 4] + [7 8] = [10 12]

[1 2]   [5 6]   [-4 -4]
[3 4] - [7 8] = [-4 -4]
```

**Properties:** commutative (`A+B = B+A`), associative
(`(A+B)+C = A+(B+C)`), additive identity `0`, additive inverse `-A`.

**Pitfalls:** dimensions must match — adding a `2 × 3` to a `3 × 2` is an
error, not a transpose. There is no "broadcasting" in the mathematical
definition (NumPy adds it as a convenience; decide explicitly whether you
want it, because it silently hides dimension bugs).

---

## 2. Scalar multiplication

**Notation:** `cA` or `c · A`

**Definition:** Multiply every entry by the scalar `c`:

```text
(cA)_ij = c · a_ij
```

**Example:**

```text
    [1 2]   [3  6]
 3· [3 4] = [9 12]
```

**Properties:** `c(A+B) = cA + cB`, `(c+d)A = cA + dA`,
`(cd)A = c(dA)`, `1·A = A`, `0·A = 0`.

**Pitfalls:** with floating point, `-1 · 0` may produce `-0`, which prints as
`-0.000` and looks like a bug. Normalize negative zero on output.

---

## 3. Transpose

**Notation:** `Aᵀ` (some texts use `A'` or `A^T`)

**Definition:** Reflect across the main diagonal: rows become columns.

```text
(Aᵀ)_ij = a_ji
```

Dimensions swap: `m × n` becomes `n × m`.

**Example:**

```text
      [1 2 3]              [1 4]
 A =  [4 5 6]    Aᵀ  =     [2 5]
                           [3 6]
```

**Properties:** `(Aᵀ)ᵀ = A`; `(A + B)ᵀ = Aᵀ + Bᵀ`; `(cA)ᵀ = cAᵀ`; and the one
worth memorizing because the order reverses: `(AB)ᵀ = BᵀAᵀ`. `A` is
**symmetric** if `A = Aᵀ` (necessarily square), **skew-symmetric** if
`A = -Aᵀ` (diagonal must be zero).

**Pitfalls:** the `(AB)ᵀ` order reversal is the most commonly misremembered
identity in the subject.

---

## 4. Trace

**Notation:** `tr(A)`

**Definition:** The sum of the diagonal entries. Defined only for square
matrices:

```text
tr(A) = Σ_i a_ii = a_11 + a_22 + ... + a_nn
```

**Example:**

```text
      [1 2]
 A =  [3 4]      tr(A) = 1 + 4 = 5
```

**Properties:** `tr(A + B) = tr(A) + tr(B)`; `tr(cA) = c·tr(A)`;
`tr(Aᵀ) = tr(A)`; and the **cyclic property** `tr(AB) = tr(BA)` (for
compatible dimensions — note `AB` and `BA` are generally different matrices,
but their traces are equal). The trace equals the sum of the eigenvalues, so
it is a cheap partial fingerprint of a matrix.

**Pitfalls:** strictly square-only. Some libraries define a "trace" for
rectangular matrices as the sum of `min(m,n)` diagonal entries; that is a
convention, not the mathematical trace, and if you support it you should say
so. For products of three or more, cyclic rotation is allowed
(`tr(ABC) = tr(BCA)`) but arbitrary reordering is not.

---

## 5. Identity matrix

**Notation:** `I_n` (or just `I`)

**Definition:** The `n × n` matrix with 1s on the main diagonal and 0s
elsewhere:

```text
I_ij = 1 if i = j, 0 otherwise
```

**Example:**

```text
        [1 0 0]
 I_3 =  [0 1 0]
        [0 0 1]
```

**Why it matters:** it is the multiplicative identity: `AI = IA = A` for any
compatible `A`. Its columns are the standard basis vectors `e_1, ..., e_n`,
and `A·e_j` extracts the `j`-th column of `A` — which is the cleanest way to
see why matrix multiplication works the way it does.

**Pitfalls:** this is the one operation in your list that takes a dimension
rather than a matrix as input, so give it a distinct input form and don't let
it silently produce a `0 × 0` matrix on an empty argument. Generating `I_n`
for large `n` is memory-linear; that is fine.

---

## 6. Matrix–vector multiplication

**Notation:** `y = Ax`

**Definition:** For `A` an `m × n` matrix and `x` an `n × 1` vector, `y` is
`m × 1`, with each entry the dot product of a row of `A` with `x`:

```text
y_i = Σ_j a_ij · x_j
```

Two equally important readings:

1. **Row reading:** `y_i` is the dot product of row `i` of `A` with `x`.
2. **Column reading:** `y` is a linear combination of the **columns** of `A`,
   where `x_j` is the coefficient on column `j`:
   `y = x_1·(col 1 of A) + x_2·(col 2 of A) + ...`

The column reading is the conceptual heart of the subject — it is what
`Ax = b` being solvable means (is `b` in the span of the columns?).

**Example:**

```text
      [1 2]           [1·2 + 2·1]   [4]
 A =  [3 4] ,  x = [2, 1]  →  y = [3·2 + 4·1] = [10]
```

**Dimension rule:** `(m × n) · (n × 1) = (m × 1)`. The inner dimensions must
agree; the outer dimensions give the result.

**Pitfalls:** the single most common user error is dimension mismatch, so
make that error message name both shapes explicitly. Cost: `2mn` flops.

---

## 7. Norms

A norm measures the "size" of a vector or matrix. Every norm satisfies:
`||x|| ≥ 0`, `||x|| = 0` iff `x = 0`, `||cx|| = |c|·||x||` (homogeneity),
`||x + y|| ≤ ||x|| + ||y||` (triangle inequality).

### Vector p-norms

```text
||x||_p = ( Σ_i |x_i|^p )^(1/p)        for p ≥ 1
```

The three that matter in practice:

| Norm | Formula | Name | Meaning |
| ---- | ------- | ---- | ------- |
| `\|\|x\|\|_1` | `Σ\|x_i\|` | Manhattan / taxicab | total magnitude |
| `\|\|x\|\|_2` | `√(Σ x_i²)` | Euclidean | ordinary geometric length |
| `\|\|x\|\|_∞` | `max_i\|x_i\|` | Chebyshev / max | largest component |

**Example:** `x = (3, -4)`

```text
||x||_1 = |3| + |-4|           = 7
||x||_2 = √(3² + (-4)²) = √25  = 5      ← the 3-4-5 triangle
||x||_∞ = max(|3|, |-4|)       = 4
```

### Matrix norms

- **Frobenius norm:** treat the matrix as a long vector of all its entries.
  `||A||_F = √(Σ_ij a_ij²) = √(tr(AᵀA))`. Cheap, always defined, dimension-agnostic.
- **Induced 1-norm:** `||A||_1 = max_j Σ_i |a_ij|` — the largest absolute
  **column** sum.
- **Induced ∞-norm:** `||A||_∞ = max_i Σ_j |a_ij|` — the largest absolute
  **row** sum.
- **Spectral (2-) norm:** `||A||_2` = largest singular value. Requires the
  SVD, so it is **out of scope** for Phase 1/2 — listed in the README's
  non-goals so nobody expects it.

**Pitfalls:** the induced norms are not symmetric in the way people expect —
the expression for `||A||_1` sums down columns while `||A||_∞` sums across
rows, which is a genuine source of bugs in hand-written implementations.
Normalizing a vector means dividing by its norm; guard against the zero
vector, since dividing by `0` is the classic nan-producer.

---

## 8. Elementwise operations (Hadamard and friends)

**Notation:** `A ∘ B` for the Hadamard (elementwise) product

**Definition:** Operate entry by entry, requiring identical dimensions:

```text
(A ∘ B)_ij = a_ij · b_ij
```

The same idea applies to elementwise division, powers (`(A^∘2)_ij = a_ij²`),
absolute value, negation, and clamping.

**Example:**

```text
[1 2]   [5 6]   [1·5 2·6]   [5  12]
[3 4] ∘ [7 8] = [3·7 4·8] = [21 32]
```

Contrast with the matrix product of the same two matrices, which is
`[[19, 22], [43, 50]]`. These are completely different operations that share
a symbol-less notation in casual writing — disambiguating them in your output
is genuinely useful to users.

**Properties:** the Hadamard product is commutative and associative, and its
identity element is the all-ones matrix `J` (not `I`). It does not care about
inner dimensions at all.

**Pitfalls:** do not call it "multiplication" anywhere in your interface. It
is the operation users most often confuse with `AB`. Cost: `n²` operations,
i.e. trivially cheap compared to the real product.

---

# Part 2 — Phase 2 operations

## 9. Matrix multiplication

**Notation:** `C = AB`

**Definition:** For `A` of size `m × n` and `B` of size `n × p`:

```text
c_ij = Σ_k a_ik · b_kj
```

Each entry of `C` is the dot product of **row `i` of `A`** with **column `j`
of `B`**. The inner dimensions must agree (`n`), and the result is `m × p`.

The two readings, mirroring the matrix–vector case:

- **Row reading:** each row of `C` is a linear combination of the rows of `B`,
  weighted by the corresponding row of `A`.
- **Column reading:** each column of `C` is a linear combination of the
  columns of `A`, weighted by the corresponding column of `B`.

The columns reading is what makes `AB` compute "apply `B`, then apply `A`":
matrix multiplication _is_ composition of linear maps, which is why the order
matters.

**Example:**

```text
      [1 2]         [5 6]
 A =  [3 4] ,  B =  [7 8]

        [1·5+2·7  1·6+2·8]   [19 22]
 A·B =  [3·5+4·7  3·6+4·8] = [43 50]

        [5·1+6·3  5·2+6·4]   [23 34]
 B·A =  [7·1+8·3  7·2+8·4] = [31 46]
```

`AB ≠ BA` — the single most important fact about the operation.

**Properties:** associative (`(AB)C = A(BC)`), distributive over addition,
`AI = IA = A`, `(AB)ᵀ = BᵀAᵀ`, `(AB)⁻¹ = B⁻¹A⁻¹` (when both invertible).
Not commutative, and not cancellative: `AB = AC` does **not** imply `B = C`.

**Pitfalls:** cost is `2mnp` flops, i.e. `O(n³)` for square matrices — the
most expensive operation in your list, and the one where a naive
implementation will visibly stall around `n ≈ 500`. Naive triple loop with
reasonable memory access order is entirely acceptable for a teaching tool;
blocking and Strassen are v2 material. Also beware integer overflow in the
accumulator if you use exact integer arithmetic.

---

## 10. Determinant

**Notation:** `det(A)` or `|A|`

**Definition:** A function from square matrices to scalars. Three equivalent
constructions:

- **Leibniz formula:** sum over all permutations `σ` of `{1..n}`, each
  weighted by the parity of the permutation:
  `det(A) = Σ_σ sgn(σ) · ∏_i a_{i,σ(i)}`.
- **Cofactor (Laplace) expansion:** expand along a row or column,
  `det(A) = Σ_j (-1)^(i+j) · a_ij · M_ij`, where `M_ij` (the minor) is the
  determinant of `A` with row `i` and column `j` deleted.
- **Via LU:** `det(A) = (sign of permutation) · ∏_i u_ii`, the product of the
  pivots from an LU factorization with row swaps. This is how you should
  compute it.

**Geometric meaning:** the factor by which the linear map scales signed
volume. In 2D it is the signed area of the image of the unit square; a zero
determinant means the map collapses space onto a lower-dimensional subspace
(hence non-invertible).

**Example (2 × 2):**

```text
      [a b]
 A =  [c d]      det(A) = ad - bc

      [1 2]
      [3 4]      det = 1·4 - 2·3 = -2
```

**Example (3 × 3):**

```text
      [1 2  3]
 A =  [4 5  6]
      [7 8 10]

 det(A) = 1·(5·10 - 6·8) - 2·(4·10 - 6·7) + 3·(4·8 - 5·7)
        = 1·(50 - 48)     - 2·(40 - 42)     + 3·(32 - 35)
        = 2 + 4 - 9 = -3
```

**Properties:** `det(AB) = det(A)·det(B)`; `det(Aᵀ) = det(A)`;
`det(cA) = cⁿ·det(A)` for `n × n` (note the exponent — a very common error);
`det(I) = 1`; `det(A) = 0` iff `A` is singular; the determinant of a
triangular matrix is the product of its diagonal; `det(A)` equals the product
of the eigenvalues.

**Pitfalls:** cofactor expansion costs `O(n!)`. For `n = 20` that is roughly
`2.4 × 10^18` operations, which will never finish — yet it is the definition
most people remember, so implement it only for tiny `n` and route everything
else through LU. In floating point, the determinant of a large matrix can
underflow to `0` or overflow to `∞` even when the matrix is perfectly
well-conditioned, because the magnitude scales like `n log n` in the exponent;
this is why real libraries compute `log|det|` or just factor instead.

---

## 11. Row reduction: REF, RREF, and Gauss-Jordan

**Notation:** `rref(A)`

**Elementary row operations** (the only three moves allowed):

1. Swap two rows: `R_i ↔ R_j`
2. Multiply a row by a nonzero scalar: `R_i ← c·R_i`, `c ≠ 0`
3. Add a multiple of one row to another: `R_i ← R_i + c·R_j`

Each corresponds to multiplying on the left by an elementary matrix, and none
of them change the solution set of a linear system or the row space.

**Row echelon form (REF):**

- All nonzero rows are above all zero rows.
- The **pivot** (leading nonzero entry) of each row is strictly to the right
  of the pivot of the row above it.
- Entries below a pivot are zero.

**Reduced row echelon form (RREF):** additionally

- Every pivot equals exactly `1`.
- Every pivot is the **only** nonzero entry in its column.

RREF is **unique** — every matrix has exactly one RREF, reachable by many
different sequences of row operations. REF is not unique.

**Gauss-Jordan elimination** is the procedure that reaches RREF: use downward
elimination to get REF, then back-substitute / eliminate upward and scale
pivots to 1. (Classical Gaussian elimination stops at REF and then
back-substitutes; Gauss-Jordan goes all the way.)

**Example:**

```text
       [1 2 1 |  4]         [1 2 0 | 3]
 A|b = [2 4 3 |  9]   →     [0 0 1 | 1]
       [3 6 4 | 13]         [0 0 0 | 0]
```

Columns 1 and 3 have pivots; column 2 does not (it is exactly 2 × column 1),
so `x_2` is **free**. Reading the rows back: `x_1 = 3 - 2·x_2`, `x_3 = 1`.
Note that pivots need not sit in adjacent columns — an elimination routine
that assumes "pivot of row `k` is in column `k`" breaks on this matrix.

**What RREF gives you for free:** rank (count the pivots), solvability (a row
`[0 0 ... 0 | c]` with `c ≠ 0` means inconsistent), the general solution
(pivot variables expressed in terms of free ones), the inverse (reduce
`[A | I]` to `[I | A⁻¹]`), and a basis for the null space.

**Pitfalls:** in exact arithmetic, choose any nonzero pivot. In floating
point, always choose the **largest-magnitude** candidate in the column
(partial pivoting) or you will amplify rounding error — the textbook
"divide by whatever is nonzero" version is numerically broken. Also decide
explicitly how you treat values that are "almost zero": an exact `== 0` test
is wrong for floats, and a loose epsilon will misclassify ill-conditioned
matrices. Make the tolerance a visible, configurable part of the tool rather
than a magic number buried in the code.

---

## 12. Rank

**Notation:** `rank(A)`

**Definition:** The number of pivots in the RREF of `A`, equivalently:

- the dimension of the **column space** (span of the columns),
- the dimension of the **row space** (span of the rows),
- the size of the largest invertible square submatrix.

For an `m × n` matrix, `rank(A) ≤ min(m, n)`. If equality holds, `A` has
**full rank**.

**Example:**

```text
      [1 2 3]
 A =  [4 5 6]        rref(A) = [1 0 -1]
      [7 8 9]                  [0 1  2]
                               [0 0  0]        →  rank(A) = 2
```

Row 3 is `2·row2 - row1`, so the rows are linearly dependent and the rank
drops to 2. Note `det(A) = 0`, consistent with `rank < n`.

**Properties:** `rank(A) = rank(Aᵀ)`; `rank(AB) ≤ min(rank(A), rank(B))`;
`rank(A + B) ≤ rank(A) + rank(B)`; `A` (square, `n × n`) is invertible iff
`rank(A) = n`.

**Rank–nullity theorem:** for an `m × n` matrix,
`rank(A) + nullity(A) = n`, where the nullity is the dimension of the null
space (the number of free variables). This theorem is the reason the three
solution cases in §14 work the way they do.

**Pitfalls:** "rank" computed by counting nonzero rows after elimination is
extremely sensitive to floating-point tolerance — singular values that are
mathematically exactly zero come out as `1e-16`. Either use exact arithmetic
or expose the tolerance and document that the answer is threshold-dependent.

---

## 13. Inverse

**Notation:** `A⁻¹`

**Definition:** For square `A`, the matrix satisfying

```text
A · A⁻¹ = A⁻¹ · A = I
```

It exists iff `det(A) ≠ 0` (equivalently `rank(A) = n`); such a matrix is
**nonsingular** or **invertible**. The inverse is unique when it exists.

**Formula (2 × 2) — worth memorizing:**

```text
      [a b]               1      [ d -b]
 A =  [c d]      A⁻¹ = ─────── · [ -c  a]
                        ad - bc
```

**Example:**

```text
      [1 2]                      1    [ 4 -2]   [-2   1 ]
 A =  [3 4]      det = -2, A⁻¹ = ── · [ -3  1] = [ 1.5 -0.5]
                                  -2
```

Verify: `A·A⁻¹ = I`. ✓

**General method:** form the augmented matrix `[A | I]` and row-reduce to
`[I | A⁻¹]`. This works for any `n` and reuses your Gauss-Jordan code, which
is why it is the natural implementation once §11 exists.

**Properties:** `(A⁻¹)⁻¹ = A`; `(AB)⁻¹ = B⁻¹A⁻¹` (order reverses, same as
transpose); `(Aᵀ)⁻¹ = (A⁻¹)ᵀ`; `det(A⁻¹) = 1/det(A)`; the inverse of two or
more factors reverses the whole order: `(ABC)⁻¹ = C⁻¹B⁻¹A⁻¹`.

**Pitfalls:** **never solve `Ax = b` by computing `x = A⁻¹b`.** It costs about
three times as much as an LU solve and produces a less accurate answer.
Computing an inverse is legitimate when you genuinely need the inverse matrix
itself, or need to solve against many right-hand sides — and even then, a
stored factorization is usually better. Also be honest about
**conditioning**: for `A` near singular, `A⁻¹` has huge entries and any
floating-point answer is unreliable; a condition-number warning is a
high-value feature.

---

## 14. Solving `Ax = b` — and the three cases

**Notation:** `Ax = b`, with `A` an `m × n` matrix, `x` the unknown `n × 1`
vector, `b` the known `m × 1` vector.

**Method:** form the **augmented matrix** `[A | b]` and reduce it. Comparing
`rank(A)` with `rank([A | b])` settles everything, because appending `b` can
raise the rank by at most one.

Let `n` be the number of unknowns (columns of `A`).

### Case 1 — unique solution

Condition: `rank(A) = rank([A | b]) = n`.

Then there are no free variables, and the solution is a single point.

```text
 x +  y = 3
 x -  y = 1

      [1  1 | 3]         [1 0 | 2]
      [1 -1 | 1]   →     [0 1 | 1]        →  x = 2, y = 1
```

### Case 2 — infinitely many solutions

Condition: `rank(A) = rank([A | b]) = r < n`.

Then there are `n - r` free variables, and the solution set is an affine
subspace of dimension `n - r`:

```text
x = x_particular + (any combination of null-space basis vectors)
```

```text
 x +  y = 2
2x + 2y = 4

      [1 1 | 2]         [1 1 | 2]
      [2 2 | 4]   →     [0 0 | 0]        →  y = t, x = 2 - t
```

### Case 3 — no solution

Condition: `rank(A) < rank([A | b])`.

The extra rank appears as a contradiction row `[0 0 ... 0 | c]` with `c ≠ 0`,
i.e. the system claims `0 = c`. Geometrically, `b` is not in the column space
of `A`.

```text
 x +  y = 2
2x + 2y = 5

      [1 1 | 2]         [1 1 | 2]
      [2 2 | 5]   →     [0 0 | 1]        →  0 = 1, inconsistent
```

### Summary table

| `rank(A)` vs `rank([A\|b])` | vs `n` | Outcome |
| --------------------------- | ------ | ------- |
| equal | `= n` | exactly one solution |
| equal | `< n` | infinitely many (`n - rank` free parameters) |
| `rank(A) < rank([A\|b])` | — | no solution |

**Square nonsingular case:** when `m = n = rank(A)`, the standard computation
is an LU factorization followed by forward and backward substitution.
Cost `≈ (2/3)n³` for the factorization, reused for every right-hand side.

**Falling outside these cases:** an overdetermined system (`m > n`) is
generically inconsistent, and the natural question becomes **least squares**
(`minimize ||Ax - b||_2`), which is a v2 candidate — it is a _different_
problem, not a fourth case.

**Pitfalls:** report which case applies explicitly. Silently returning
`x = 0` for an inconsistent system, or a single arbitrary point for a family
of solutions, is the most user-hostile failure mode this tool can have. For
Case 2, present the general solution with the free variables named (`y = t`),
not one numeric witness.

---

## 15. LU factorization with pivoting

**Notation:** `PA = LU`

**Definition:** Factor `A` into

- `L`: lower triangular with 1s on the diagonal (unit lower triangular),
- `U`: upper triangular,
- `P`: a permutation matrix recording the row swaps.

Written `PA = LU`, or equivalently `A = PᵀLU` since permutation matrices are
orthogonal (`Pᵀ = P⁻¹`). Without row swaps you get `A = LU`, which fails
whenever a pivot is zero — and is numerically dangerous whenever a pivot is
merely small.

**Example (why pivoting is not optional):**

```text
      [0 2]
 A =  [1 3]
```

Naive elimination needs to divide by `a_11 = 0`. Swap the rows
(`P = [[0,1],[1,0]]`):

```text
        [1 3]                [1 0]        [1 3]
 P·A =  [0 2]      L =       [0 1] ,  U = [0 2]      →  PA = LU ✓
```

`det(A) = sgn(P)·u_11·u_22 = (-1)·1·2 = -2`, matching `0·3 - 2·1 = -2`. ✓

**Partial pivoting (the standard strategy):** at step `k`, look at the
entries in column `k` at or below the diagonal and swap the
**largest-magnitude** one into the pivot position. This keeps all multipliers
`|l_ik| ≤ 1`, which bounds the growth of rounding error. It does not require
the pivot to be nonzero — it _guarantees_ it will be, unless the matrix is
singular.

**Algorithm sketch (Doolittle, with partial pivoting):**

```text
U ← A, L ← I, P ← I
for k = 1..n-1:
  find row r ≥ k maximizing |U[r][k]|;
  swap rows r and k in U, L (cols < k), and P
  if U[k][k] == 0: matrix is singular, stop
  for i = k+1..n:
      L[i][k] = U[i][k] / U[k][k]
      U[i][k:] -= L[i][k] · U[k][k:]
```

**Uses:**

- Solving `Ax = b`: factor once, then `Pb = LUx` becomes a forward
  substitution `Ly = Pb` and a backward substitution `Ux = y`. Reusing the
  factorization for many `b` is the whole point.
- Determinant: product of `U`'s diagonal, times the sign of the permutation.
- Inverse: solve against each standard basis vector `e_j`.
- Rank and linear independence, indirectly, via the pivot structure.

**Cost:** about `(2/3)n³` flops for the factorization. A full solve is
`2n³/3 + O(n²)`, versus `2n³` for forming `A⁻¹` and then multiplying.

**Pitfalls:** in exact arithmetic, a row swap is needed only when the pivot
candidate is zero, and any nonzero pivot is correct. In floating point,
partial pivoting is what keeps the result stable. Partial pivoting has no
theoretical guarantee against element growth (complete pivoting, which also
swaps columns, does), but it is stable in practice for essentially all inputs
you will meet. For rank and RREF on singular or rectangular matrices, a column
with no nonzero candidate is skipped, not an error — the "singular, stop" line
in the sketch above applies only to LU of a square matrix. Watch the permutation sign tracking:
it is easy to lose a factor of `-1` in the determinant, and that error is
invisible in a solve but wrong in every displayed determinant.

---

# Part 3 — Adjacent operations worth reconsidering

These are not on the current roadmap but are cheap, conceptually central, and
arguably belong in Phase 1. Worth a deliberate decision rather than an
oversight.

| Operation | Definition | Why it might belong |
| --------- | ---------- | ------------------- |
| Vector addition / subtraction / scaling | Same as §1–2 with `n × 1` operands | Users will reach for these first; they're a few lines given the matrix code |
| Dot product | `u·v = Σ u_i v_i = uᵀv` | The atom that matrix–vector multiply is built from; useful on its own and shows up in projections and angles |
| Outer product | `u vᵀ`, an `m × n` matrix of rank 1 | The mirror image of the dot product, and the cleanest illustration of rank |
| Matrix power | `A^k = A·A···A` (`k` times) | Trivial once §9 exists; naturally raises "what about non-integer `k`?" which motivates eigenvalues |
| Submatrix / slicing | Extract rows/columns/block | Needed internally by cofactor expansion and LU; exposing it makes the tool scriptable |
| Concatenation / stacking | `[A \| B]` horizontally, `[A ; B]` vertically | The augmented matrix of §14 is exactly horizontal concatenation; exposing it lets users reproduce the solve by hand |
| Row operations as first-class commands | Swap, scale, `R_i ← R_i + cR_j` | These are the steps the Phase 2 step-by-step output will display; exposing them turns the tool into a teaching aid |
| Diagonal extraction / construction | `diag(A)` ↔ `diag(v)` | Cheap, used everywhere, and the bridge to the trace |
| `log \| det(A) \|` | Avoids underflow/overflow of the raw determinant | The numerically honest way to report a determinant |

---

# Part 4 — Out of scope, potentially v2

Grouped by theme. These are the operations a linear algebra tool is
_expected_ to grow into, which is also why the README's non-goals list matters: they
are all defensible, and all expensive.

## 4.1 Eigen-theory

| Operation | What it is | Notes |
| --------- | ---------- | ----- |
| Eigenvalues / eigenvectors | `Av = λv` with `v ≠ 0`: vectors the map only stretches | The single biggest missing feature; the gateway to everything below |
| Characteristic polynomial      | `p(λ) = det(A - λI)`; its roots are the eigenvalues                    | Requires symbolic or polynomial arithmetic to do properly                     |
| Diagonalization                | `A = PDP⁻¹`, `D` diagonal of eigenvalues                               | Possible iff there are `n` independent eigenvectors                           |
| Similarity                     | `B = P⁻¹AP`; same eigenvalues, different basis                         | The reason diagonalization is useful                                          |
| Spectral theorem               | Every real symmetric matrix is orthogonally diagonalizable: `A = QDQᵀ` | The nicest special case; underlies PCA and least squares                      |
| Jordan canonical form          | The best you can do when a matrix is not diagonalizable                | Genuinely hard; almost certainly never                                        |
| Positive/negative definiteness | `xᵀAx > 0` (resp. `< 0`) for all `x ≠ 0`                               | Testable via eigenvalues or leading principal minors; central to optimization |
| Quadratic forms                | `q(x) = xᵀAx`                                                          | The scalar function a symmetric matrix represents; conic classification       |
| Singular values & SVD          | `A = UΣVᵀ`; the `σ_i` are the square roots of `AᵀA`'s eigenvalues      | The most important factorization in applied mathematics; also unlocks `\|\|A\|\|_2`, pseudoinverse, low-rank approximation, PCA |
| Pseudoinverse                  | `A⁺ = VΣ⁺Uᵀ`, the best generalized inverse                             | Solves least squares and rank-deficient systems in one framework              |
| Condition number               | `κ(A) = \|\|A\|\| · \|\|A⁻¹\|\|`; in the 2-norm, `σ_max/σ_min`       | How much relative error the solve amplifies; a warning the tool should raise |
| Power iteration                | Repeatedly apply `A` and normalize to get the dominant eigenpair       | The simplest iterative eigenvalue algorithm; a good self-contained project    |
| QR algorithm                   | Iterate `A_k = Q_kR_k`, `A_{k+1} = R_kQ_k`                             | How real libraries compute all eigenvalues; hard                              |
| Rayleigh quotient              | `xᵀAx / xᵀx`, an eigenvalue estimate for unit `x`                      | Cheap, elegant, pairs well with power iteration                               |
| Deflation                      | Remove a known eigenpair to find the next one                          | The standard way to get more than the dominant eigenvalue                     |

## 4.2 Factorizations

| Factorization        | Form                                              | Notes                                                                                                                                                            |
| -------------------- | ------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| QR                   | `A = QR`, `Q` orthogonal, `R` upper triangular    | Via Gram-Schmidt (unstable as written), Householder reflections (stable, standard), or Givens rotations (sparse-friendly); unlocks least squares and eigenvalues |
| Cholesky             | `A = LLᵀ` for symmetric positive definite `A`     | Half the cost of LU and no pivoting needed; the numerical gold standard for SPD systems                                                                          |
| `LDLᵀ`               | Pivoted variant for symmetric indefinite matrices | Sits between Cholesky and LU                                                                                                                                     |
| Schur                | `A = QTQᵀ`, `T` upper triangular                  | Every square matrix has one; the complex-number-free route to eigenvalues                                                                                        |
| Hessenberg reduction | Orthogonal similarity to upper Hessenberg form    | The preprocessing step before the QR algorithm                                                                                                                   |
| Polar decomposition  | `A = QP` with `Q` orthogonal, `P` PSD             | The matrix analogue of `z = re^{iθ}`; used in mechanics and graphics                                                                                             |
| Rank factorization   | `A = CF` with `C` full column rank                | Exposes the row and column spaces explicitly                                                                                                                     |
| Kronecker product    | `A ⊗ B`, a block matrix whose blocks are `a_ij·B` | The standard tool for vectorizing matrix equations                                                                                                               |

## 4.3 Solving and approximation

| Operation                       | What it is                                                         | Notes                                                                   |
| ------------------------------- | ------------------------------------------------------------------ | ----------------------------------------------------------------------- |
| Least squares                   | `minimize \|\|Ax - b\|\|₂`; the normal equations `AᵀAx = Aᵀb`      | The right answer for overdetermined systems; solve via QR rather than the normal equations, which square the condition number |
| Over/underdetermined systems    | `m > n` or `m < n`, generically inconsistent or under-constrained  | A natural extension of §14 once least squares exists                    |
| Cramer's rule                   | `x_i = det(A_i)/det(A)`, where `A_i` replaces column `i` with `b`  | Worth _documenting_ for its elegance, never worth implementing: `O(n!)` |
| Gauss–Seidel / Jacobi iteration | Iterative solvers that converge to `x` for diagonally dominant `A` | Cheap to implement, pedagogically interesting, poor convergence         |
| Conjugate gradient              | Krylov method for symmetric positive definite systems              | The go-to iterative solver; needs only matrix–vector products           |
| GMRES                           | Krylov method for general sparse systems                           | The non-symmetric counterpart to CG                                     |
| Preconditioning                 | Transform a system to improve conditioning before iterating        | What makes iterative solvers actually fast                              |
| Matrix functions                | `exp(A) = Σ A^k/k!`, `log(A)`, `√A`                                | Reachable from diagonalization; genuinely tricky in general             |

## 4.4 Structure, scale, and theory

| Topic                                                 | Why it matters                                                                                                         |
| ----------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| Null space, column space, row space, basis, dimension | The conceptual vocabulary behind rank and solvability; a tool could print these directly                               |
| Orthogonality, projections, Gram–Schmidt              | The geometry underlying QR and least squares                                                                           |
| Sparse matrices (`CSR`/`CSC` formats)                 | Real-world matrices are mostly zeros; unlocks large problems                                                           |
| Banded, triangular, Toeplitz, circulant structure     | Special-cased algorithms that are dramatically faster                                                                  |
| Block matrices and Schur complements                  | Compositional tricks for structured problems                                                                           |
| Vectorization (`vec` operator) and its identities     | `vec(AXB) = (Bᵀ ⊗ A)vec(X)`; makes matrix calculus mechanical                                                          |
| Matrix calculus                                       | Derivatives with respect to vectors and matrices; the language of machine learning                                     |
| Numerical stability, backward error, error analysis   | Why the "obvious" algorithm is often the wrong one — arguably more valuable than any single operation here             |
| Exact / rational / symbolic arithmetic                | Adopted as the default (see README → Design decisions); makes determinants, RREF, and ranks _unambiguous_ rather than tolerance-dependent |
| Complex numbers                                       | Needed for non-symmetric eigenvalues and the Schur form                                                                |
| Parallel and GPU implementations                      | Where `O(n³)` actually bites                                                                                           |
| BLAS / LAPACK bindings                                | The pragmatic route to performance, and the reason a from-scratch tool needs a stated purpose                          |

---

# Part 5 — Numerics notes that apply throughout

1. **Exact vs. floating point changes your answers, not just their precision.**
   With `float64`, `det` of an integer matrix can come back as
   `-5.999999999999998`, `rank` depends on a tolerance, and `rref` may produce
   `1e-17` where a zero belongs. With exact rationals, these problems vanish
   and the results are verifiable by hand. The tradeoff is range and speed.
2. **Choosing a pivot is a numerical decision, not a notational one.** Always
   select the largest-magnitude candidate in the column.
3. **Never test floats with `==`.** For zero tests, use a documented,
   configurable tolerance — and make it visible in the output when it changes
   a result.
4. **Don't square the condition number.** Normal equations (`AᵀAx = Aᵀb`) are
   the classic trap in least squares; QR-based methods avoid it.
5. **Cost table** (square `n × n`, leading term):

   | Operation                                    | Flops                                      |
   | -------------------------------------------- | ------------------------------------------ |
   | `A + B`, `A ∘ B`, scalar multiply, transpose | `n²`                                       |
   | `Ax`                                         | `2n²`                                      |
   | `AB` (naive)                                 | `2n³`                                      |
   | LU factorization                             | `≈ (2/3)n³`                                |
   | Solve `Ax = b` given `LU`                    | `≈ 2n²`                                    |
   | Solve `Ax = b` from scratch                  | `≈ (2/3)n³`                                |
   | Inverse via `[A \| I]`                       | `≈ 2n³` (and it's the less accurate route) |
   | Determinant by cofactor expansion            | `O(n!)` — do not                           |

6. **Display and computation are different problems.** Computing a matrix
   product is `O(n³)`; rendering it attractively with aligned columns is
   `O(n²)` string work, and in practice it is where the bugs will be
   (negative signs, decimals, color codes counting toward column width,
   Unicode brackets being double-width in some terminals). Budget for it
   accordingly.
