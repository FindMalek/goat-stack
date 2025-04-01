# Testing

How we’ve configured testing in Goat Stack.

Goat Stack uses Vitest for unit testing. It’s a fast and lightweight testing framework that’s compatible with Jest’s API. Unit tests are run automatically as part of the build task in Turborepo.

## Running Tests

To run tests, simply execute:

```bash
pnpm test
```

This will run all tests in the `__tests__` folder in each app, however we’ve only written a couple of tests for app project so far.

## Adding Tests

We use Testing Library for our tests. It’s a great library that allows you to test your components in a way that’s close to how a user would interact with them.

In the `__tests__` folder, create a new file called `[page].test.tsx` (where `[page]` is the name of the page you want to test). Here’s an example of a test for the home page:

`apps/app/__tests__/home.test.tsx`

```typescript
import { render, screen } from '@testing-library/react';
import { expect, test } from 'vitest';
import Page from '../app/(unauthenticated)/home/page';

test('Home Page', () => {
  render(<Page />);
  expect(
    screen.getByRole('heading', {
      level: 1,
      name: 'Hello, world!',
    })
  ).toBeDefined();
});
```

## Adding Vitest to other apps

To add Vitest to another app, you’ll need to install the dependencies:

```bash
pnpm install -D vitest @testing-library/react @testing-library/dom --filter [app]
```

as well as adding the `@goat/testing` package to the `devDependencies` section of the `package.json` file:

`apps/[app]/package.json`

```json
"devDependencies": {
  "@goat/testing": "workspace:*"
}
```

Create a new file called `vitest.config.ts` in the root of the app and add the following:

`apps/[app]/vitest.config.ts`

```typescript
export { default } from '@goat/testing';
```

Then, create a new file in the `__tests__` folder in the relevant app and add a test command to the `package.json` file:

`apps/[app]/package.json`

```json
{
  "scripts": {
    "test": "vitest"
  }
}
```

Turborepo will automatically pick up on the new test script and run the tests.

Then, just follow the instructions above to write tests.
