An App is a sub-module of a project.

It is used to implement functionality for some purpose.

Apps can be self-contained, so it does not rely on other apps to function.

This allows them to be re-used in different projects.

A project can contain many apps.

E.g. 

You have a social media project:
Then you have an app for the news feed.
An app for the comments
An app for the posts
An app for the friends list

Apps are added using the `startapp` command:

````
python manage.py startapp feed
python manage.py startapp comments 
python manage.py startapp profile 
``
