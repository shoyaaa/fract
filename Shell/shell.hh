#ifndef __SHELL_HH
#define __SHELL_HH

#include <iostream>

#include "../Objects/color.hh"
#include "../Utilities/string.hh"
#include "../Include/operating_system.hh"

/// @brief Mark of pwd line.
#define pwd_mark ">"

namespace Fract::Shell {
///  @brief Utilities for shell.
class shell {
public:
  /**
   * @brief Print message with color.
   * @param msg Message.
   * @param color Color of message.
   */
  static void printc(std::string msg, std::string color);

  /**
   * @brief Print error.
   * @param msg Error message.
   */
  static void print_error(std::string msg);

  /**
   * @brief Get input with message and color.
   * @param msg Message.
   * @param color Color.
   * @return Input.
   */
  static std::string get_input(std::string msg, std::string color);

  /**
   * @brief Get input with message.
   * @param msg Message.
   * @return Input.
   */
  static std::string get_input(std::string msg);

  /**
   * @brief Get input with pwd.
   * @return Input.
   */
  static std::string get_input();
};
} // namespace Fract::Shell

#endif // __SHELL_HH
