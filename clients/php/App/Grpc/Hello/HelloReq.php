<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto

namespace App\Grpc\Hello;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * message 对应生成代码的 struct
 * 定义客户端请求的数据格式
 * &#64;validator=HelloReq
 *
 * Generated from protobuf message <code>App.Grpc.Hello.HelloReq</code>
 */
class HelloReq extends \Google\Protobuf\Internal\Message
{
    /**
     * [修饰符] 类型 字段名 = 标识符;
     * &#64;inject_tag: json:"id" validate:"required,min=1"
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    protected $id = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *           [修饰符] 类型 字段名 = 标识符;
     *           &#64;inject_tag: json:"id" validate:"required,min=1"
     * }
     */
    public function __construct($data = NULL) {
        \App\Grpc\GPBMetadata\Hello::initOnce();
        parent::__construct($data);
    }

    /**
     * [修饰符] 类型 字段名 = 标识符;
     * &#64;inject_tag: json:"id" validate:"required,min=1"
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * [修饰符] 类型 字段名 = 标识符;
     * &#64;inject_tag: json:"id" validate:"required,min=1"
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

}

