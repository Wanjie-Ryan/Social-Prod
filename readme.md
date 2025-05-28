--> go mod init github.com/Wanjie-Ryan/Social-Prod --> Standard go intialization, Go will expect the project to be hosted at the Github URL if it were to be published
--> This way, Go can fetch dependencies from Github directly.
--> bin folder will contain the compiled Go code.
--> cmd folder will have the executables or entry point for the app CMD ---> api (any server related code will be here, http, middleware, etc)
--> internal will have the internal packages not to be exported
--> docs will contain swagger related code
--> scripts

***** PRINCIPLES ****
1. Separation of concerns
-- Each level in your program should be separate by a clear barrier.
2. Dependency Inversion Principle
-- Inject dependencies in your layers, you don't directly call them.
3. Adapatablitity to change
-- Organize code in a modular and flexible way, so as it can be easy to introduce new features.
4. Business value
-- Focus on delivering value to your users.