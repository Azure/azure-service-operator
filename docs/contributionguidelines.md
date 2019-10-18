# Contribution Workflow Guidelines

1. You should use a [fork-and-branch](https://blog.scottlowe.org/2015/01/27/using-fork-branch-git-workflow/) strategy for contribution to this repo.

2. During development, we will base PRs off a feature branch in the main fork, instead of master until an operator is ready. For eg, we will merge all SQL operator PRs into the branch "azure-sql" and not master.

3. If you have changes from others as part of your PR, use "Merge commit" to merge changes instead of "Squash and Merge" and it helps retain their contribution history.

4. Wait for the pipeline checks to succeed before merging your approved PR, to ensure tests pass.

5. When reviewing a PR, checkout that branch and test all relevant positive/negative cases before you approve it. Do not skip the tests before approval even if time is short.

6. We recommend smaller, more regular but logical PRs instead of a large one. It helps get feedback early and reduces chances of merge conflicts.

7. If there are changes that require redeploying the operator, pause to think if it would impact any existing deployments and indicate this in the PR.

8. If there are comments against your PR, make sure you either mark them as "Resolved" if addressed or comment if you disagree. It is not acceptable to re-submit a PR without acknowledging a comment.

9. If there is someone who requests changes on your PR, please re-request review from them to make sure they get to review the updates.

10. If you have a PR that has more than 25 commits for comments on the PR, it might be better to recreate the PR after syncing with the branch. This would help ensure the PR is more up-to-date with the branch.

11. If there are check-ins made against your PR, make sure to do a "git pull" on your PR branch again before you test and re-send the PR.
