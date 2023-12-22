from RestrictedPython import safe_builtins, utility_builtins
from RestrictedPython.Eval import default_guarded_getiter
from RestrictedPython.Guards import guarded_iter_unpack_sequence, safer_getattr
import json
import pandas as pd
import sklearn
import numpy as np
from datetime import datetime

def custom_get_item(obj, key):
    return obj[key]

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

def custom_import(name, globals=None, locals=None, fromlist=(), level=0):
    allowed_modules = ['sklearn.', 'functools']
    
    # Check if the import name starts with any of the allowed modules
    if any(name.startswith(module) or name == module.rstrip('.') for module in allowed_modules):
        return __import__(name, globals, locals, fromlist, level)
    
    raise ImportError(f"Import of '{name}' is not allowed")

# Setup the restricted environment
builtins = {}
builtins['__import__'] = custom_import
builtins['getattr'] = safer_getattr
builtins['_getattr_'] = safer_getattr
builtins["_setattr_"] = setattr
builtins["setattr"] = setattr
builtins["_getitem_"] = custom_get_item
builtins['_getiter_'] = default_guarded_getiter
builtins['_iter_unpack_sequence_'] = guarded_iter_unpack_sequence
builtins["_setitem_"] = safe_setitem
builtins["_write_"] = custom_write
builtins.update(safe_builtins)
builtins.update(utility_builtins)

safe_globals = dict(__builtins__=builtins)
safe_globals['json'] = json
safe_globals['pd'] = pd
safe_globals['sklearn'] = sklearn
safe_globals['np'] = np
safe_globals['datetime'] = datetime