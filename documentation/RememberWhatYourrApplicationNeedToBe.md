[ ] Must have automated tests (Unit, integration and E2E), use the concept of mutations in tests and try to perform coverage

[ ] It must have continuous integration and delivery processes. Execution build, automated tests, security tests, vulnerabilities, dependency checking, load tests, etc. You can create pre-commit triggers with actions, to check interesting things like code checking, formatting, security checks and commit message pattern

[ ] Must have good observability and monitoring, both in persisted data and in logs. Store information such as access (date, location, etc.), errors, performance, availability, etc.

[ ] It must be safe! Define and activate the protection of the main development and production branches, disabling '_force pushing_', the functionality of removing the branch and the execution of actions before merges, such as running tests, building, etc. GitHub has several security options, try to activate them all.

- **Threat analysis and risk modeling**: Identify potential threats to your system and assess associated risks to prioritize security efforts.

- **Security requirements**: Include specific security requirements early in the development process, covering aspects like authentication, authorization, encryption, and access control.

- **Team training**: Ensure the development team is aware of security best practices and common vulnerabilities through regular training.

- **Code reviews and security testing**: Conduct code reviews to identify and fix security vulnerabilities before deployment. Perform regular security testing, such as penetration testing, to identify and address any existing vulnerabilities.

- **Principle of least privilege**: Adopt the principle of least privilege, granting users and system components only the access they need to perform their functions.

- **Continuous monitoring**: Implement continuous monitoring tools to detect suspicious activities or potential security breaches in real-time.

- **Incident response**: Develop and test an incident response plan to effectively respond to security breaches and minimize their impact.

[ ] Must be optimized and scalable

[ ] It must be readable, sustainable and organized

[ ] It must have good documentation, a good _Readme_, about versioning (_changelog.md_), routes, components, architecture, database (migrations), etc. Use the concept of semantic versioning in sets with releases/tags. Have in drawing (draw.io) the general system context of the application, with external dependencies, entry and exit points. Also have in private documentation the trust levels, access permissions, categorization (Stride) and risk rating (Dread) of possible threats, which is the most valuable thing in the application

[ ] Updated: Maintenance of dependencies and updating versions, keep project dependencies up to date, including libraries, frameworks and plugins used. This is important to maintain security and compatibility

[ ] Must be aware of data protection laws (Brazil LGPD), attribute the necessary credits to external sources, inform the user about _copyright_, GDPR (General Data Protection Regulation) or other applicable regulations. Have a _LICENSE_ file in the project.

[ ] User education: It is important to educate users about good security practices, such as the importance of choosing strong passwords, not reusing passwords across multiple services, and being aware of phishing attempts.

## Just for Back-End

[ ] Use standard middleware/interceptors to add common behavior to all requests that pass through the application. Check if the user is authenticated (without repeating code), error monitoring, tracking, recording metrics, etc.

[ ] Use standard guards to check whether a given user has permission to access a given component

[ ] Use component to standardize success or error returns

## Just for Front-End

[ ] **Must be accessible and inclusive for everyone**. Dont forget the accessibility tests.

[ ] **Must be responsive**.

[ ] [Reset css](https://www.alura.com.br/artigos/o-que-e-reset-css).

[ ] adding normalize (node_modules)

[ ] Use _BEM_ for CSS naming standard

[ ] Global css files: '_\_breakpoints_' file for responsiveness and '_\_variables_' file to save colors, spacing, etc.

- **Mobile**: 768
- **Desktop_xsm**: 980
- **Desktop_sm**: 1080
- **Desktop_md**: 1280

[ ] [Set global font](https://fonts.google.com/)

[ ] Set a global background, primary and secondary colors: [ColorHunt](https://colorhunt.co/) || [Color Adobe](https://color.adobe.com/en/create/color-wheel)

[ ] Set a **maximum** and **minimum** screen size for your site to 'work responsive'

```css
html {
  min-width: 320px;
  max-width: 1400px;
}
```

[ ] Your project must be internationalizable

[ ] Have a good SEO (Search Engine Optimization) plan. Create the briefing documentation file (objectives for each part of the application), pay attention to the project title, project description, names, descriptions, buttons and images of the application (put alt for Google to search), etc.

[ ] It must have basic components such as _footer_ with copyright, pages such as 'About' and 'Contact Us', privacy policy, terms of use, etc. It is also important to get permissions to display ads with _ADSense_

[ ] Have error-friendly components:

- Log errors with console.log to know where they are in the prod
- Catch error splitting during rendering phase or other lifecycle. There are APIs for handling errors like _react-error-boundary_.

❌ Use inappropriate tone

❌ Use technical jargon

❌ Pass the blame

❌ Be generic

✅ Tell us what happened and why

✅ Provide security

✅ Be empathetic

✅ Help them fix it

✅ Give them an exit, like a contact link for support or a “try again” button

[ ] There must be an option between _Dark-Mode_ and _White-Mode_:

- Check **user preference in the first color selected when opening the application**

[ ] Implement _Theming_, remembering that it is different from _Dark-Mode_ and _White-Mode_, an example, we can have the _neon_ theme that has a light mode and a dark mode... [Video on YouTube tutorial on TailWind](https:/ /www.youtube.com/watch?v=TavBrPEqkbY&ab_channel=SkiesDev)

[ ] Be creative on page 204 _notfound_
