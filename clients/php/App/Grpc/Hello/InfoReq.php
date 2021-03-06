<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto

namespace App\Grpc\Hello;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * &#64;validator=InfoReq
 *
 * Generated from protobuf message <code>App.Grpc.Hello.InfoReq</code>
 */
class InfoReq extends \Google\Protobuf\Internal\Message
{
    /**
     * 主要用于grpc validator参数校验
     * &#64;inject_tag: json:"name" validate:"required,min=1"
     *
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           主要用于grpc validator参数校验
     *           &#64;inject_tag: json:"name" validate:"required,min=1"
     * }
     */
    public function __construct($data = NULL) {
        \App\Grpc\GPBMetadata\Hello::initOnce();
        parent::__construct($data);
    }

    /**
     * 主要用于grpc validator参数校验
     * &#64;inject_tag: json:"name" validate:"required,min=1"
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * 主要用于grpc validator参数校验
     * &#64;inject_tag: json:"name" validate:"required,min=1"
     *
     * Generated from protobuf field <code>string name = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

}

