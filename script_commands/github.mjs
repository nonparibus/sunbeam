#!/usr/bin/env zx

// @raycast.schemaVersion 1
// @raycast.title List My Repositories
// @raycast.mode command
// @raycast.packageName Github

$.verbose = false;

if (argv.issues) {
  const out = await $`gh issue list --repo ${argv.issues} --json title,url`;
  const issues = JSON.parse(out);
  console.log(
    JSON.stringify({
      type: "list",
      list: {
        items: issues.map((issue) => ({
          title: issue.title,
          actions: [
            {
              type: "open-url",
              title: "Open in Browser",
              url: issue.url,
            },
          ],
        })),
      },
    })
  );
  process.exit(0);
}

var res
if (argv.owner) {
  res = await $`gh repo list ${argv.owner} --json nameWithOwner,description,url`;
} else {
  res = await $`gh repo list --json nameWithOwner,description,url`;
}
const repos = JSON.parse(res);

console.log(
  JSON.stringify({
    type: "list",
    list: {
      items: repos.map((repo) => ({
        title: repo.nameWithOwner,
        subtitle: repo.description,
        actions: [
          {
            type: "callback",
            title: "List Issues",
            args: ["--issues", repo.nameWithOwner],
          },
          {
            type: "open-url",
            title: "Open in Browser",
            url: repo.url,
          },
        ],
      })),
    },
  })
);
