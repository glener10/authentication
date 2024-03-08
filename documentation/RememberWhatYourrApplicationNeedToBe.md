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
