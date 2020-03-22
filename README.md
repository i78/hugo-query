# Hugo-Query
Hugo-Query extracts content from a Hugo content folder structure.

## Fields of application
hugo-query is supposed to be used in release pipeline steps for extracting content from the raw .md markdown files in applications where it is not feasible to obtain the wanted data from another source.

# Examples

## Extract eMail addresses from a user list
I created hugo-query to extract a whitelist of email addresses from a team list, with one page per individual containing a members email address in a frontmatter field named "email". The whitelist is generated in a pipeline step before the hugo public folder is uploaded to the target server.

Given a tree of Markdown files with a header named "email":

```
+++
title = "Alice Bob"
type = "member"
country = "norway"
email = "sendmeallyourspam@example.com"
+++
Lorem Ipsum dolor set amet ..
```

.. one can extract the email field as a json map as follows:


```
hugo-query extract --content-folder my-hug-site/content --content-type member --field email --pretty
```

Output:

```
{
  "Alice Bob": {
    "email": "sendmeallyourspam@example.com"
  }
}
```

It is also possible to extract multiple fields:
```
hugo-query extract --content-folder my-hug-site/content --content-type member --field email --field country --pretty
```

Output:

```
{
  "Alice Bob": {
    "email": "sendmeallyourspam@example.com",
    "country": "norway"
  }
}
```

## Code of Conduct
This project adheres to No Code of Conduct.  We are all adults.  We accept anyone's contributions.  Nothing else matters.

For more information please visit the [No Code of Conduct](https://github.com/domgetter/NCoC) homepage.


