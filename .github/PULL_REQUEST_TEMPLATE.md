#### Issue Info

[![Jira](https://img.shields.io/badge/Jira-0052CC?style=for-the-badge&logo=Jira&logoColor=white)](https://minnek.atlassian.net/browse/${ticket})

### Code Review Checklist for Authors

- [ ] Update your branch with the latest changes from `develop`.
- [ ] Verify all the `checks` are passed (if applicable).
- [ ] Use [code conventions](https://en.wikipedia.org/wiki/Coding_conventions) and [best practices](https://en.wikipedia.org/wiki/Best_practice).
- [ ] New features must have compatibility for a11y (Accessibility).
- [ ] No unintentional `console.logs` left behind after debugging.
- [ ] Notify the code reviewers on time to have an efficient review time.

#### Shopify Checklist

- [ ] Linted with [Theme Check](https://github.com/Shopify/theme-check)
- [ ] Tested on [mobile](https://shopify.dev/themes/store/requirements#mobile-browser-requirements)
- [ ] Tested on [multiple browsers](https://shopify.dev/themes/store/requirements#desktop-browser-requirements)
- [ ] Tested for [accessibility](https://shopify.dev/themes/best-practices/accessibility)

### Code Review Checklist for Reviewers

- [ ] Take some time to understand the code you are reading.
- [ ] Use an inquisitive tone, do not make an order.
- [ ] Accept that many programming decisions are opinions. Engage a discussion and reach a resolution quickly.
- [ ] Seek to understand the author’s perspective.

>**In case something is not applied, justify the reason why you skip one of the points above**

#### Notes before Merge

- Change pull request title to `Feature/Ticket-ID` to match the branch name. See the docs [here](https://github.com/Minnek-Digital-Studio/minnek-developer-handbook/blob/master/development/git-version-control.md#pull-request).
- Write Short and Detailed Commit Messages. See the docs [here](https://github.com/Minnek-Digital-Studio/minnek-developer-handbook/blob/master/development/git-version-control.md#writing-commits).
- Set a `developer`, to Code Review this pull request before merge to the `develop` branch. See the docs [here](https://github.com/Minnek-Digital-Studio/minnek-developer-handbook/blob/master/development/code-review.md).
- Follow these Guides for Coding Standard: [SASS](https://github.com/bigcommerce/sass-style-guide), [JavaScript](https://developer.mozilla.org/en-US/docs/MDN/Guidelines/Code_guidelines/JavaScript) and [HTML5](https://developer.mozilla.org/en-US/docs/MDN/Guidelines/Code_guidelines/HTML#class_and_id_names).
