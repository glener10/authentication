Route to send a code by email and save in database in code_verify_email
Removing  jwt tokens for REST routes
Route to verify email, checking the code with code_verify_email in database

Password Recovery: use a unique token sended to email
Update change email to receive a unique code sended for email

Gateway and method to send email and put in 'update change email', 'password recovery' and 'email verification' usecases
CI/CD Securely (Snyk, SonarQube)
2FA
Login with google

Improvement feature: Notify user when a strange login occurs, necessary to save information such as IP address, geographic location and device used. When this happens, it is necessary to save the information from the new login device and send an email asking the user. If it wasn't, it should reset all known devices and inactivate the user
Improvement feature: Common middleware to block operation from another user if doesnt admin: BlockOperationFromAnotherUserIfNotAdmin, removing Jwt Middleware because just check signature and another middlewares already due this