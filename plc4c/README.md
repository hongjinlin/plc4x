<!--
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
  -->

## Setting up in CLion

Per default CLion will not be able to run the tests. 
However it's pretty simple to setup.

In order to make this happen, you should add another `Profile` to your `CMake` configuration.

You can do this by going to: 

`Preferences...` / `Build, Execution, Deployment` / `CMake`

In the `Profiles` list, click on the `+` button to add a new profile.

I gave this profile the name `Test`.

In the settings select the `Build type` = `Debug`, `Toolchain` = `Use Default`.
`CMake options` = `-DUNITY_VERSION:STRING=2.5.2 -DBUILD_PHASE=test-compile`.

Leave the rest unchanged (which is actually empty).

After saving, you can select the `Test` profile in the `Configurations` Drop-Down.
After that is selected, the tests should be available.

If you like to switch between building the full project with IntelliJ and the PLC4X part in CLion, it is recommendable to configure CLion to place the CMake related stuff in the `target` directory.
You can do this by opening the `Settings...` dialog, then in `Build, Execution, Deployment ...` select `CMake` and in `Build directory` simply enter `target`.
This way the settings are not located outside the target directory and interfere with the maven build.
The only thing you have to keep in mind, is that if you did a maven build, you need to select `Reload CMake Project` on the `CMakeList.txt` file in the `plc4c` root directory in order to recreate the CLion configuration.
