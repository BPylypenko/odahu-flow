"""
Test models
"""

import drun.model
import drun.types as types
import drun.io

import unittest2
import pandas
import numpy


class TestScipyModel(unittest2.TestCase):
    """
    Unit tests for models
    """

    def test_positive(self):
        def prepare(x):
            x['additional'] = x['d_int']
            return x

        def apply(x):
            assert type(x) == pandas.DataFrame

            assert x['d_int'].dtype == numpy.int
            assert x['d_float'].dtype == numpy.float
            assert x['d_str'].dtype in [numpy.str, numpy.object]
            assert x['additional'].dtype == numpy.int

        df = pandas.DataFrame([{
            'd_int': 1,
            'd_float': 1.0,
            'd_str': 'what?'
        }])

        s = drun.model.ScipyModel(
            apply,
            prepare,
            drun.io._get_column_types(df),
            version='1.0')

        s.apply({'d_int': '1', 'd_float': '2.0', 'd_str': 'omg'})

    def test_typecast(self):

        df = pandas.DataFrame([{
            'd_int': 1,
            'd_float': 1.0,
            'd_str': 'what?',
            'excessive': False
        }])

        class CustomBoolObject(types.BaseType):
            def __init__(self):
                super(CustomBoolObject, self).__init__(default_numpy_type=numpy.bool_)

            def parse(self, value):
                """
                Parse boolean strings like 'of course', 'not sure'
                :param value:
                :return:
                """
                str_value = value.lower()
                if str_value == 'of course':
                    return True
                elif str_value == 'not sure':
                    return False
                else:
                    raise Exception('Invalid value: %s' % (str_value,))

            def export(self, value):
                return value

        def apply(x):
            assert type(x) == pandas.DataFrame

            assert x['d_int'].dtype == numpy.int
            assert x['d_float'].dtype == numpy.float
            assert x['d_str'].dtype in [numpy.str, numpy.object]
            assert x['excessive'].dtype == numpy.bool_

        s = drun.model.ScipyModel(
            apply,
            lambda x: x,
            column_types=drun.io._get_column_types((df, {'excessive': CustomBoolObject()})),
            version='1.0')

        s.apply({'d_int': '1', 'd_float': '2.0', 'd_str': 'omg', 'excessive': 'of course'})
        print(s.description)


if __name__ == '__main__':
    unittest2.main()