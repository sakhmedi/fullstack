# Strategy Scenarios

## Trunk-Based Development

### Description
In Trunk-Based Development, all developers commit directly to the main branch.
There are no long-lived feature branches — work is small, frequent, and continuously integrated.

### Example
- Commit message: "Trunk-Based: Add hello world function"
- Screenshot shows commit going directly into main branch.

### Advantages
- Continuous integration, fewer merge conflicts.
- Very fast feedback cycle.

### Disadvantages
- Requires strict discipline and constant code reviews.
- Mistakes can affect the main branch immediately.

---

## GitHub Flow

### Description
GitHub Flow uses short-lived feature branches.  
Each new feature or fix is made in a separate branch, followed by a Pull Request to `main`.

### Example
- Branch: `feature-add-goodbye`
- Commit message: "GitHub Flow: Add goodbye function"
- Screenshot shows PR creation and merge.

### Advantages
- Code review and discussion before merge.
- Safe to experiment without breaking main.

### Disadvantages
- Slightly slower process.
- Requires manual merge.

---

## My Opinion

I find **GitHub Flow** more convenient because it allows code review and safer collaboration.  
It’s easier to manage features and keep the main branch stable while still working quickly.
