
#include "soup.h"
#include "tofu_soup.h"

soup_t * tofu_soup_new()
{
	soup_t * soup = soup_new();
	soup->name = "Tofu Soup";  // this is FAIL
	soup->ingredients = arraylist_string_new();
	arraylist_string_add( soup->ingredients,"1 Pound tofu");
	arraylist_string_add( soup->ingredients,"1 cup carrot juice");    
	arraylist_string_add( soup->ingredients,"1/4 cup spirolena");         
	return soup;
}

