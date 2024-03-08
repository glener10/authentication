[ ] It must have tests and constantly execute them in an automated way, routines and continuous integration procedures.

[ ] Must have good observability and monitoring, both in persisted data and in logs. Store information such as access (date, location, etc.), errors, performance, availability, etc.

[ ] It must be safe! Define and activate the protection of the main development and production branches, disabling '_force pushing_', the functionality of removing the branch and the execution of actions before merges, such as running tests, building, etc. GitHub has several security options, try to activate them all.

[ ] Must be optimized and scalable

[ ] It must be readable, sustainable and organized

[ ] It must have good documentation, a good _Readme_, about versioning (_changelog.md_), routes, components, architecture, database (migrations), etc. Use the concept of semantic versioning in sets with releases/tags.

[ ] Updated: Maintenance of dependencies and updating versions, keep project dependencies up to date, including libraries, frameworks and plugins used. This is important to maintain security and compatibility

[ ] Must be aware of data protection laws (Brazil LGPD), attribute the necessary credits to external sources, inform the user about _copyright_, GDPR (General Data Protection Regulation) or other applicable regulations

[ ] Use standard request interceptors to add common behavior to all requests that pass through the application. These interceptors allow you to intercept requests before they are handled by controllers or services and perform additional actions, such as error monitoring, tracking, recording metrics, authentication, validation, among others.

[ ] Use standard exception filters to always return the same structure

[ ] User education: It is important to educate users about good security practices, such as the importance of choosing strong passwords, not reusing passwords across multiple services, and being aware of phishing attempts.

## If is Front-End

[ ] **Must be accessible and inclusive for everyone**.

[ ] **Must be responsive**.

[ ] Your project must be internationalizable (Example react-i18next library).

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

[ ] Must be optimized

- Use the **developer tool** (browser extension) to see which components render more than once (appear with a green box on the screen)

- Use lazy loading/code division (**lazy loading**). Only load what is needed by the user, as needed, load the rest.

- **Always import specifically and not a whole package from a library!**

[ ] There must be an option between _Dark-Mode_ and _White-Mode_:

- Check **user preference in the first color selected when opening the application**

[ ] Implement _Theming_, remembering that it is different from _Dark-Mode_ and _White-Mode_, an example, we can have the _neon_ theme that has a light mode and a dark mode... [Video on YouTube tutorial on TailWind](https:/ /www.youtube.com/watch?v=TavBrPEqkbY&ab_channel=SkiesDev)

[ ] Be creative on page 204 _notfound_