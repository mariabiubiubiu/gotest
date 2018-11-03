#
# Autogenerated by Thrift Compiler (0.10.0)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py
#

from thrift.Thrift import TType, TMessageType, TFrozenDict, TException, TApplicationException
from thrift.protocol.TProtocol import TProtocolException
import sys

from thrift.transport import TTransport


class Student(object):
    """
    Attributes:
     - sid
     - sname
     - ssex
     - sage
    """

    thrift_spec = (
        None,  # 0
        (1, TType.I32, 'sid', None, None, ),  # 1
        (2, TType.STRING, 'sname', 'UTF8', None, ),  # 2
        (3, TType.BOOL, 'ssex', None, False, ),  # 3
        (4, TType.I16, 'sage', None, None, ),  # 4
    )

    def __init__(self, sid=None, sname=None, ssex=thrift_spec[3][4], sage=None,):
        self.sid = sid
        self.sname = sname
        self.ssex = ssex
        self.sage = sage

    def read(self, iprot):
        if iprot._fast_decode is not None and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None:
            iprot._fast_decode(self, iprot, (self.__class__, self.thrift_spec))
            return
        iprot.readStructBegin()
        while True:
            (fname, ftype, fid) = iprot.readFieldBegin()
            if ftype == TType.STOP:
                break
            if fid == 1:
                if ftype == TType.I32:
                    self.sid = iprot.readI32()
                else:
                    iprot.skip(ftype)
            elif fid == 2:
                if ftype == TType.STRING:
                    self.sname = iprot.readString().decode('utf-8') if sys.version_info[0] == 2 else iprot.readString()
                else:
                    iprot.skip(ftype)
            elif fid == 3:
                if ftype == TType.BOOL:
                    self.ssex = iprot.readBool()
                else:
                    iprot.skip(ftype)
            elif fid == 4:
                if ftype == TType.I16:
                    self.sage = iprot.readI16()
                else:
                    iprot.skip(ftype)
            else:
                iprot.skip(ftype)
            iprot.readFieldEnd()
        iprot.readStructEnd()

    def write(self, oprot):
        if oprot._fast_encode is not None and self.thrift_spec is not None:
            oprot.trans.write(oprot._fast_encode(self, (self.__class__, self.thrift_spec)))
            return
        oprot.writeStructBegin('Student')
        if self.sid is not None:
            oprot.writeFieldBegin('sid', TType.I32, 1)
            oprot.writeI32(self.sid)
            oprot.writeFieldEnd()
        if self.sname is not None:
            oprot.writeFieldBegin('sname', TType.STRING, 2)
            oprot.writeString(self.sname.encode('utf-8') if sys.version_info[0] == 2 else self.sname)
            oprot.writeFieldEnd()
        if self.ssex is not None:
            oprot.writeFieldBegin('ssex', TType.BOOL, 3)
            oprot.writeBool(self.ssex)
            oprot.writeFieldEnd()
        if self.sage is not None:
            oprot.writeFieldBegin('sage', TType.I16, 4)
            oprot.writeI16(self.sage)
            oprot.writeFieldEnd()
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        return

    def __repr__(self):
        L = ['%s=%r' % (key, value)
             for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)