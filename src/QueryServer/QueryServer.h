/**
 * Tencent is pleased to support the open source community by making Tars available.
 *
 * Copyright (C) 2016THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed 
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR 
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the 
 * specific language governing permissions and limitations under the License.
 */

#ifndef _QueryServer_H_
#define _QueryServer_H_

#include <iostream>
#include <set>
#include "servant/Application.h"

using namespace std;
using namespace tars;

/////////////////////////////////////////////////////////////////////
class QueryServer : public Application {
public:
    /**
     *
     **/
    ~QueryServer() override = default;;

    /**
     *
     **/
    void initialize() override;

    /**
     *
     **/
    void destroyApp() override;

    const std::string &getELKIndexPre() const;

    void getELKNodeAddress(std::string &host, int &port);

private:
    std::mutex mutex;
    std::string elkIndexPre{};
    std::vector<std::tuple<std::string, int>> elkTupleNodes{};
};

extern QueryServer g_app;
////////////////////////////////////////////
#endif
