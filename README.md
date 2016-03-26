![Liberty CLI](http://i.imgur.com/JkcYy4c.png)

*liberty - the state of being free within society from oppressive restrictions imposed by authority on one's behaviour or political views*


### Note
This project is currently a proposal, it's not ready for use yet and depends heavily on those willing to contribute. If you're interested in helping, please don't hesitate.


### Yet another dependency manager?

The motivation behind Liberty comes from the lack of independence that current solutions such as npm provide. We do not have any control over
developers code. We respect everybody.

A recent situation occurred involving npm and one of their respected contributors. Ultimately, he was treated poorly and after the situation was resolved there was absolutely no apology from npm.

This has developers screaming out for a new package manager. One that didn't have complete control over your code. One that respected developers over their business model and corporate companies. The best way to do that? Don't have a business model.


### What is Liberty

Liberty is a simple, friendly and powerful dependency manager that has no single point of failure. We do this by avoiding a central list of packages and we go straight to the source.

There is a plan to have a website where you can search for "libs". This will be achieved by using data in the "liberty" file that developers provide. We will never replicate your code without your explicit permission. Your code belongs to you.

Liberty is proud to offer:
- A simple, convenient way to manage dependencies
- No single point of failure
- No namespace, no conflicts


### Example Configuration
```
{
	“liberty” :
	{
		“author” : “Swagger Developer”,
		“title” : “Dank Project”,
		“desc” : “A dank, swagoo package manager”
	},
	“dependencies” :
	{
		“github:otherdeveloper/lib-dependency” : “1.3.0”,
		“github:otherdeveloper/bin-dependency” : “1.4.0a”
	}
}
```
