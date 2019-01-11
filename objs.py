#!/usr/bin/env python2

class obj:
    def __init__(self): pass

    def obj_method(self):
        print("Python object method called!")
        return 1

    def obj_method_arg(self, arg1):
        print("Python object method called, with argument: %s" % arg1)
        return 1

    def obj_method_args(self, arg1, arg2):
        print("Python object method called with arguments: %s" % arg1, arg2)
        return 1

    def obj_method_kwargs(self, **kwargs, args):
        print("Python object method called with arguments:" )
        print(kwargs)
        return 1
#
# if __name__ == '__main__':
#     obj_a = obj()
#     obj_a.obj_method_kwargs(**{'key': 'val'})