// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// see the license for the specific language governing permissions and
// limitations under the license.
#ifndef CXX_SPEC_MOCK_ANALYZER_H_
#define CXX_SPEC_MOCK_ANALYZER_H_

#include "gmock/gmock.h"
#include "cxx/spec/analyzer.h"

namespace mako {
class MockAnalyzer : public Analyzer {
 public:
  MOCK_METHOD(bool, ConstructHistoricQuery,
              (const mako::AnalyzerHistoricQueryInput& query_input,
               mako::AnalyzerHistoricQueryOutput*),
              (override));
  MOCK_METHOD(bool, Analyze,
              (const mako::AnalyzerInput& analyzer_input,
               mako::AnalyzerOutput*),
              (override));
  MOCK_METHOD(std::string, analyzer_type, (), (override));
  MOCK_METHOD(std::string, analyzer_name, (), (override));

 private:
  bool DoAnalyze(const mako::AnalyzerInput& analyzer_input,
                 mako::AnalyzerOutput* analyzer_output) override {
    return true;
  }
};
}  // namespace mako

#endif  // CXX_SPEC_MOCK_ANALYZER_H_
