#!/usr/bin/python

import time, dialog, sys
from types import *



class Result:

  _exception = None
  _value = None
  _success = False

  @staticmethod
  def from_exception( exc ):
    if exc is None: raise Exception( 'Argument "exc" must not be None.' )
    if type( exc ) is not ExceptionType: raise Exception( 'Argument "exc" must be of type Exception.' )
    r = Result()
    r._exception = exc
    return r

  @staticmethod
  def from_value( val ):
    if val is None: raise Exception( 'Argument "val" must not be None.' )
    r = Result()
    r._value = val
    r._success = True
    return r

  @staticmethod
  def from_failure():
    return Result()

  def exc( self ):
    return self._exception
  
  def val( self ):
    return self._value
 
  def success( self ):
    return self._success

  def failure( self ):
    return not self._success

  def __str__( self ):
    return repr( self._value ) + ", " + repr( self._exception ) + ", " + str( self._success )



class Chooser:
 
  import dialog

  def __init__( self, title, choices ):

    if title is None: self._title = 'Choose'
    else: self._title = title

    if choices is None: raise Exception( 'Argument "choices" must not be None.' )
    self._choices = []

    for c in choices:
      if type( c ) is not TupleType: raise Exception( 'Argument is not a tuple.', repr( c ) )
      if len( c ) != 3: raise Exception( 'Tuple must have 3 elements.', repr( c ) )
      if c[ 0 ] is None: raise Exception( 'First argument must not be None.' )
      if type( c[ 0 ] ) not in ( IntType, StringType ): raise Exception( 'First argument must be an integer or a string.' )
      if c[ 1 ] is None: raise Exception( 'Second argument must not be None.' )
      if type( c[ 1 ] ) is not StringType: raise Exception( 'Second argument must be a string.' )
      if c[ 2 ] is None: raise Exception( 'Third argument must not be None.' )
      if type( c[ 2 ] ) is not StringType: raise Exception( 'Third argument must be a string.' )

      c = ( str( c[ 0 ] )[ 0:1 ] + ' ' + c[ 1 ], c[ 2 ] )
      self._choices.append( c )

    self._dialog = dialog.Dialog( dialog = 'dialog' )


  def choose( self ):

    ( code, tag ) = self._dialog.menu( self._title,
					width = 60,
					choices = self._choices )
    if code == 0:
      return Result.from_value( tag )
    return Result.from_failure()

        

    
 

c = Chooser( 'ssh to DV', [ 
  ( 1, 'dv-jb1', 'DV JBoss Server 1' ),
  ( 2, 'dv-jb2', 'DV JBoss Server 2' ),
  ( 3, 'dv-jb3', 'DV JBoss Server 3' ),
  ] )
print c.choose()
sys.exit( 0 )



class Actions:

  def do1( self ):
    print 1

  def do2( self ):
    print 2

  def do3( self ):
    print 3



actions = Actions()

dia = dialog.Dialog( dialog = 'dialog' )
print dia



while True:
  ( code, tag ) = dia.menu( 
    'ssh to',
    width = 60, 
    choices = [ ( '1 anton', 'anton' ),
		( '2 berta', 'bbsjkjdsbj' ),
		( '3 cciji', 'cjidejoidjeoidjeijdejdie' ),
	      ] )
  break

if code == 0:
  sel = tag[ 0:1 ]
  getattr( actions, 'do' + sel )()
  
  
  

