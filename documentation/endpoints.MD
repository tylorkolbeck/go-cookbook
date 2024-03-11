POST /cookbooks
Allows a user to create a new cookbook. The request body should include details like the cookbook's name, description, and initial visibility status (public or private).

GET /cookbooks
Retrieves a list of all public cookbooks. Can be extended with query parameters to filter by user, name, etc., for enhanced search functionality. For authenticated users, this endpoint could also list the user’s private cookbooks.

GET /cookbooks/{cookbookId}
Returns detailed information about a specific cookbook, including its recipes. Access control should be implemented to ensure private cookbooks are only viewable by their owner or users with explicit access.

PUT /cookbooks/{cookbookId}
Allows updates to cookbook properties (e.g., name, description, visibility). Restricted to the cookbook's owner or users with edit permissions.

DELETE /cookbooks/{cookbookId}
Permanently removes a cookbook. This action should be heavily guarded with permissions to prevent accidental deletions.

POST /recipes
Allows a user to create a new standalone recipe. The request body should include the recipe's details such as title, ingredients, and instructions.

GET /recipes
Retrieves a list of all public recipes. Can be extended with query parameters to filter by user, ingredients, etc., for enhanced search functionality. For authenticated users, this could also list their private recipes.

GET /recipes{receipeId}
Returns detailed information about a specific recipe. Access control should be implemented to ensure private recipes are only viewable by their creator or users with explicit access.

PUT /recipes{receipeId}
Allows updates to recipe properties. Restricted to the recipe's creator.

DELETE /recipes{receipeId}
Permanently removes a standalone recipe. Restricted to the recipe's creator.
