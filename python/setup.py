from RestrictedPython import compile_restricted
from RestrictedPython.Guards import safe_builtins
import json
import pandas as pd

def custom_get_item(obj, key):
    return obj[key]

def custom_get_iter(obj):
    return iter(obj)

def safe_setitem(obj, key, value):
    # Implement your logic to safely set items
    # For example, you might want to check if obj is a DataFrame and key is allowed
    obj[key] = value
    return obj

def custom_write(obj):
	"""
	Custom hooks which controls the way objects/lists/tuples/dicts behave in
	RestrictedPython
	"""
	return obj

# custom_safe_builtins = (
# 	'sorted',
# 	'reversed',
# 	'map',
# 	# 'reduce',
# 	'any',
# 	'all',
# 	'slice',
# 	'filter',
# 	'len',
# 	'next',
# 	'enumerate',
# 	'sum',
# 	'abs',
# 	'min',
# 	'max',
# 	'round',
# 	# 'cmp',
# 	'divmod',
# 	'str',
# 	# 'unicode',
# 	'int',
# 	'float',
# 	'complex',
# 	'tuple',
# 	'set',
# 	'list',
# 	'dict',
# 	'bool',
# )

# Setup the restricted environment
builtins = safe_builtins.copy()
builtins["_getattr_"] = getattr
builtins["getattr"] = getattr
builtins["_setattr_"] = setattr
builtins["setattr"] = setattr
builtins["_getitem_"] = custom_get_item
builtins["_getiter_"] = custom_get_iter
builtins["_setitem_"] = safe_setitem
builtins["_write_"] = custom_write
builtins["len"] = len
builtins["range"] = range
builtins["list"] = list
builtins["dict"] = dict
builtins["print"] = print  # Add print to the builtins

# Layer in our own additional set of builtins that we have
# considered safe.
# for key in custom_safe_builtins:
# 	builtins[key] = __builtins__[key]

safe_globals = dict(__builtins__=builtins)
safe_globals['json'] = json
safe_globals['pd'] = pd