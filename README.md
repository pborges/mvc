# MVC example

this is very much a work in progress!

##Highlights
* Call server side methods on models without javascript.
* not much else ha

##Concepts
* mvc.ViewModel has all the data required to draw a page (using mvc.RenderLayout)
    * mvc.ViewModel.ViewBag is a concept borrowed from C# MVC I thought was useful
* mvc.Form is used to draw a postback form for a struct (and is usually used as mvc.ViewModel.Model)
    * this struct is used to give the form to allow multiple models (and sub models) in a form post# mvc_example
